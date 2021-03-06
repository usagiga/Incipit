# Docker

We can use Docker & Docker Compose to set up Incipit.


## Run Incipit on Docker Compose

Run it,

```sh
$ docker-compose up
```

then Docker Compose builds and runs Incipit automatically.


## For Developers

If you want to run test or use debugger on virtual environments, you can use special `docker-compose-*.yml` .

- `docker-compose-test.yml`
    - It runs `go test` once
- `docker-compose-debug.yml`
    - It runs Incipit through [go-delve/delve](https://github.com/go-delve/delve/) and [cosmtrek/air](https://github.com/cosmtrek/air/)
        - Hot reload enabled when editing `*.go` or `*.tmpl`
        - delve runs on `:2345`
    - Optimal for Go Remote on GoLand. See also [GoLand blog](https://blog.jetbrains.com/go/2019/02/06/debugging-with-goland-getting-started/#debugging-a-running-application-on-a-remote-machine)
