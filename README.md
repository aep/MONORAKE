THE MONORAKE
------------

static website builder that just works

 - scss
 - candybars templates (handlebars + lua)
 - cache busting
 - zero dependencies, it will still work tomorrow
 - really really fast


download binaries here: https://github.com/aep/MONORAKE/releases


### processing

put all your stuff in a directory named source/ in whatever structure.
the monorake processes those files according to their file endings from left to right.

so for a file named **source/bla/thingy.candy.scss.hash.css**

 1. candybars
 2. scss
 3. hashed
 4. copied to dist/bla/thingy-87128973198273.css

available processors:

 - scss:   https://sass-lang.com/
 - candy:  https://mustache.github.io/mustache.5.html but with lua
 - nop:    does nothing (see processing passes)
 - hash:   append hash of file content to file name
 - layout: wrap in layout.candy.html

## link to hashed file

hashes are added to filenames for cache busting, so browsers re-fetch it when the contents change.
to link the file from html or whatever, a global is available in candybars

given a file src/img/lul.hash.jpeg, ref it like this:

```html
    <img src="{{ref_img_lul_jpeg}}" alt="lol">
```

don't worry about figuring out the right variable name, it is printed during processing


## caveat: processing passes

there's no dependency tree, because i've got more important things to do, so processing is done in passes.
Given two files with identical processing depth, like doo.html.candy and dah.html.candy,
one cannot refer to the other as hashed file.

you can work around it by adding a nop processor like dah.html.candy.nop which delays
that file to one pass later than doo.html.candy

if someone actually cares i might implement it correctly, but until then this works well enough for me
