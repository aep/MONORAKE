package src

import (
    "github.com/pelletier/go-toml"
    "io/ioutil"
    "os"
    "io"
    "strings"
    "fmt"
    "html/template"
    "path/filepath"
    "crypto/sha256"
    "encoding/hex"

)


type Config struct {
}

type State struct {
     Config         Config

     forwardmap     map[string]string
     reversemap     map[string]string
     Path           func (i string) string
}

var state State;

func Path (i string) string {
    i2, ok := state.forwardmap[i]
    if !ok {
        panic("cannot find source file '"+ i + "'");
    }
    i = i2
    for ;; {
        i2, ok := state.forwardmap[i]
        if !ok {
            return i;
        }
        i = i2;
    }
}




func Main() {
    fmt.Println("~!~ THE MONORAKE ~!~");

    tstr, err := ioutil.ReadFile("monorake.toml")
    if err != nil {
        panic(err);
    }

    state = State {
        forwardmap: make(map[string]string),
        reversemap: make(map[string]string),
        Path:   Path,
    }



    err = toml.Unmarshal(tstr, &state.Config)
    if err != nil { panic(err);}





    err = os.RemoveAll("dist")
    if err != nil { panic(err);}
    err = os.MkdirAll("dist", os.ModePerm)
    if err != nil { panic(err);}



    err = filepath.Walk("src", func(path string, info os.FileInfo, err error) error {

        if info == nil {
            return nil
        }

        dst := "dist/" + strings.Join(strings.Split(path, "/")[1:], "/")


        if info.IsDir() {
            err := os.MkdirAll(dst, os.ModePerm)
            if err != nil { panic(err) }
            return nil
        }

        src, err := os.Open(path)
        if err != nil { panic(err) }
        defer src.Close()

        destFile, err := os.Create(dst)
        if err != nil { panic(err) }
        defer destFile.Close()

        _, err = io.Copy(destFile, src)
        if err != nil { panic(err) }

        err = destFile.Sync()
        if err != nil { panic(err) }

        return nil
    });
    if err != nil { panic(err) }




    for i:=1;;i++ {

        settled := true

        fmt.Println("\n:: iteration", i);



        err := filepath.Walk("dist", func(path string, info os.FileInfo, err error) error {

            if info.IsDir() {
                return nil
            }
            exts := strings.Split(path, ".");
            if len(exts) < 3 {
                return nil
            }

            src, err := os.Open(path)
            if err != nil { panic(err) }

            dstpath := exts[0] + "." + strings.Join(exts[2:len(exts)], ".")
            dst, err := os.Create(dstpath)
            if err != nil { panic(err) }


            switch exts[1] {
                case "nop":
                    fmt.Println("   nop     " , path);
                    _, err = io.Copy(dst, src)
                    if err != nil { panic(err) }
                    break;
                case "tpl":
                    fmt.Println("   template " , dstpath);
                    templates, err := template.ParseFiles(src.Name())
                    if err != nil { fmt.Println(""); panic(err) }
                    err = templates.ExecuteTemplate(dst, info.Name(), &state)
                    if err != nil { fmt.Println(""); panic(err) }
                    break;
                case "layout":
                    templatename := "layout.html"
                    fmt.Println("   layout " , dstpath, "<<", templatename);
                    templates, err := template.ParseFiles(templatename, src.Name())
                    if err != nil { fmt.Println(""); panic(err) }
                    err = templates.ExecuteTemplate(dst, "layout.html", &state)
                    if err != nil { fmt.Println(""); panic(err) }
                    break;
                case "hash":

                    _, err = io.Copy(dst, src)
                    if err != nil { panic(err) }
                    src.Seek(0,0)
                    hash := sha256.New()
                    _, err = io.Copy(hash, src)
                    if err != nil { panic(err) }

                    hashhex := hex.EncodeToString(hash.Sum(nil))

                    fmt.Println("   hash    " , path, hashhex);


                    dstpath = exts[0] + "-"+ hashhex + "." + strings.Join(exts[2:len(exts)], ".")

                    state.forwardmap[strings.Join(strings.Split(exts[0], "/")[1:], "/") + "." + strings.Join(exts[2:len(exts)], ".")]= 
                        strings.Join(strings.Split(dstpath, "/")[1:], "/")

                    err = os.Rename(dst.Name(), dstpath)
                    if err != nil { panic(err) }


                default:
                    panic(path + ": unknown filter: " + exts[1])

            }

            state.forwardmap[strings.Join(strings.Split(path, "/")[1:], "/")] = strings.Join(strings.Split(dstpath, "/")[1:], "/")
            state.reversemap[strings.Join(strings.Split(dstpath, "/")[1:], "/")]  = strings.Join(strings.Split(path, "/")[1:], "/")

            dst.Sync()
            dst.Close()
            src.Close()
            os.Remove(src.Name())

            settled = false

            return nil
        })
        if err != nil { panic(err) }

        if settled {
            break
        }
    }
}



func layout() {
}

