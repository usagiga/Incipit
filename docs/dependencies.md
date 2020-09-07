# Dependencies

Please set it up before using Incipit in your own server.

- Go (1.14 or higher)
- MariaDB (10.4 or higher)
- npm


# Develop Dependencies

If you want to set up develop environment, See [Installation](./install.md).
Or you can use Docker, See also [Docker](./docker.md).


## Backend

- Go
    - Gin
    - Gorm
    - See also: `back/go.mod`
- MariaDB
- (Optional) [cosmtrek/air](https://github.com/cosmtrek/air)
- (Optional) [go-delve/delve](https://github.com/go-delve/delve)


## Frontend

- npm
    - Nuxt.js
        - TypeScript
        - Vuetify.js
        - nuxt-property-decorator
    - ESLint
    - See also: `front/package.json`

If you use Docker, `node_modules` is NOT shared with a host machine.
So, to use ESLint or completions on IDE, please run `npm i` on your host machine.

## Infrastructure

- Docker
    - Docker Compose
- GitHub Actions
- (Optional) [h2o/h2o](https://github.com/h2o/h2o)
