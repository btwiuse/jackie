# Jackie Chain

[![Go 1.18+](https://img.shields.io/github/go-mod/go-version/btwiuse/jackie)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/btwiuse/jackie?color=%23000&style=flat-round)](https://github.com/btwiuse/jackie/blob/master/LICENSE)
[![DockerHub](https://img.shields.io/docker/pulls/btwiuse/jackie.svg)](https://hub.docker.com/r/btwiuse/jackie)
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/btwiuse/jackie)

Jack of all trades, master of some.

![Jackie Chain](https://www.movieplus.jp/film_img/CS-0000000200800677-000_l.jpg)

## Run

- using `docker`

    ```
    $ docker run -it -v $PWD/conf:/home/xuper/conf -v $PWD/data:/home/xuper/data -p 8085:8085 -p 37101:37101 btwiuse/jackie
    ```

    after that you should able to run ./demo/counter and ./demo/nft to verify it's really working. You may have to install [yj](https://github.com/sclevine/yj/releases/tag/v5.1.0) and jq first in order to run those scripts.

- using `go run`

    ```
    # add scripts to your PATH
    $ ./link

    # start the local chain node
    $ ./start

    # start the api server
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

- [/docs](https://jackie-chain.herokuapp.com/docs)
- [/redoc](https://jackie-chain.herokuapp.com/redoc)
- [/rapidoc](https://jackie-chain.herokuapp.com/rapidoc)

