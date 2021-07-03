THE MONORAKE
------------

static website builder that just works

 - scss
 - go templates
 - cache busting
 - zero dependencies, it will still work tomorrow
 - really really fast


download binaries here: https://github.com/aep/MONORAKE/releases


### processing

put all your stuff in a directory named source/ in whatever structure.
the monorake processes those files according to their file endings from left to right.

so for a file named **source/bla/thingy.template.scss.hash.css**

 1. tpl
 2. scss
 3. hashed
 4. copied to dist/bla/thingy-87128973198273.css

available processors:

 - scss:   https://sass-lang.com/
 - tpl:    https://golang.org/pkg/html/template/
 - nop:    does nothing (see processing passes)
 - hash:   append hash of file content to file name
 - layout: wrap in layout.candy.html

## link to hashed file

hashes are added to filenames for cache busting, so browsers re-fetch it when the contents change.
to link the file from html or whatever, a global is available in candybars

given a file src/img/lul.hash.jpeg, ref it like this:

```html
    <img src="{{call .Path "img/lul.hash.jpeg"}}" alt="lol">
```

## caveat: processing passes

there's no dependency tree, because i've got more important things to do, so processing is done in passes.
Given two files with identical processing depth, like doo.html.candy and dah.html.candy,
one cannot refer to the other as hashed file.

you can work around it by adding a nop processor like dah.html.candy.nop which delays
that file to one pass later than doo.html.candy

if someone actually cares i might implement it correctly, but until then this works well enough for me
