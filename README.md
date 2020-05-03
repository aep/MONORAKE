THE MONORAKE
------------

static website builder that just works

 - scss
 - candybars templates (handlebars + lua)
 - cache busting
 - zero dependencies, it will still work tomorrow
 - really really fast

### site structure

make a website like this:
```
  src/
    pages/
      index.html
      whatever.html
    layouts/
      default.html
    css/
      0_bootstrap.min.css
      1_unbreak_bootstrap.scss
      2_thingy.css
    js/
      1_spam.js
      2_spam_must_blink_hard.js
    resources/
      memes.jpeg

```


pages contains .. pages that will be baked to static html.


## candybars

is like handlebars, but with lua, so uuh, you just put any lua expression in {{HERE}} and guchi


## cache busting

It adds hashes to filesnames everywhere so browsers re-fetch it when the contents change.

given a file src/resources/neva-giev-yu-up.jpeg, this is how you use it:

```html
    <img src="{{res_neva_giev_yu_up_jpeg}}" alt="lol">
```



