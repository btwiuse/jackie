# Jackie Chain

Jack of all trades, master of some.

![Jackie Chain](https://www.movieplus.jp/film_img/CS-0000000200800677-000_l.jpg)

## Run

- using `docker`

```
$ docker run -it -p 8085:8085 btwiuse/jackie
```

- using `go run`

```
$ go run .
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

2022/05/11 18:01:22 Now you can visit http://127.0.0.1:8085/docs or http://127.0.0.1:8085/redoc to see the api docs. Have fun!
[GIN-debug] POST   /api/v1/xuper/keypair/new --> github.com/long2ice/swagin/router.IAPI.Handler-fm (3 handlers)
[GIN-debug] GET    /health                   --> github.com/long2ice/swagin/router.IAPI.Handler-fm (3 handlers)
[GIN-debug] GET    /openapi.json             --> github.com/long2ice/swagin.(*SwaGin).init.func1 (2 handlers)
[GIN-debug] GET    /docs                     --> github.com/long2ice/swagin.(*SwaGin).init.func2 (2 handlers)
[GIN-debug] GET    /redoc                    --> github.com/long2ice/swagin.(*SwaGin).init.func3 (2 handlers)
[GIN-debug] GET    /rapidoc                  --> github.com/long2ice/swagin.(*SwaGin).init.func4 (2 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
```

## Docs

- /docs
- /redoc
- /rapidoc

