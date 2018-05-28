# go-session

Small sample Go application that demonstrates how to set and delete HTTP Cookies as part of an `http.Handler`. Also how to test `http.Handler` types in general.

This application is ready to deploy to Heroku if you want to test it. It will just need some environment variables set up:

- `COOKIE_NAME`: the name of the cookie you want to set
- `COOKIE_VALUE`: the value of the cookie
- `COOKIE_DOMAIN`: the domain of the cookie
- `COOKIE_PATH`: the path of the cookie
- `COOKIE_DURATION`: how many hours you want the cookie to be valid for

To set the cookie, visit the `/login` URL, and to delete it, visit the `/logout` URL.

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/brafales/go-session)
