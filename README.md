THE MONORAKE
------------

static website builder that just works


 - scss
 - candybars templates (handlebars + lua)
 - cache busting
 - zero dependencies, it will still work tomorrow
 - really really fast

## how to use

make a structure like this:
```
  src/
    pages/
      index.html
      whatever.pug
    layouts/
      default.pug
    css/
      0_bootstrap.min.css
      1_unbreak_bootstrap.scss
      2_thingy.css
    js/
      0_spam.js
      1_morespam.js
    resources/
      memes.jpeg

```


then just bake the page by calling, you guessed it 'monorake'
dist/ now contains your page. yey
