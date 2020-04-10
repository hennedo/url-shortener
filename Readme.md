A simple URL Shortener
======================

This is the code for the url shortener running at https://dre.li - it supports emoji urls 💁‍♂️

## Requirements

  * A running MongoDB server
  
## Config

The following config flags can be used:
| Flag             | Environment Variable | Default                   | Description                      |
|------------------|----------------------|---------------------------|----------------------------------|
| --mongodb        | MONGODB              | "localhost/url-shortener" | MongoDB Connection String        |
| --port           | PORT                 | 8000                      | Port on which the server listens |
| --admin-password | ADMIN_PASSWORD       | foobar2342                | Password for the /admin endpoint |
| --base-url       | BASE_URL             | "http://localhost:8000"   | Base URL for shorturls           |

## Docker

Dockerfile is enclosed, a premade container will be available soon

## Shorten URLs!

Do it either with the Webinterface located at the root, or

  * POST / with x-www-form-urlencoded value "url" set
  * POST / with query parameter ?url=

There is a second parameter, "emoji" if you set this to "1", it will create a shorturl using emojis.


## Using it with other Software!

It is compatible with Software like dropshare. As default the shortener will assume a human is using this and render a HTML Page. 
If you are using the shortener together with other software you might want a different output. The shortener can render as text/html (default), application/json (a JSON representation of the shorturl) or text/plain (just the shorturl)

Example JSON output:
```json
{
  "shorturl":"https://dre.li/4G4zKU",
  "url":"https://google.de",
  "manageurl":"https://dre.li/4G4zKU/xxxxx"
}
```

Just set the Accept Header accordingly. If for whatever reason it is not possible for you to set the Header correct, you can set the correct encoding in the "accept" value of either the Post Form or the Query Parameter.
