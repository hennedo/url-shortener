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

Also there is a second parameter, "emoji" if you set this to "1", it will create a shorturl using emojis

It is compatible with Software like dropshare
