# Dependencies

Please set it up before using Incipit in your own server.

- Go (1.14 or higher)
- MariaDB (10.4 or higher)
- npm


# Develop Dependencies

If you want to set up develop environment, See [Installation](./install.md).


## Backend

- Go
    - Gin
    - Gorm
    - See also: `back/go.mod`
- MariaDB
- (Optional) [cosmtrek/air](https://github.com/cosmtrek/air)


## Frontend

- npm
    - Nuxt.js
        - TypeScript
        - Vuetify.js
        - nuxt-property-decorator
    - ESLint
    - See also: `front/package.json`


## Infrastructure

- Docker
    - Docker Compose
- GitHub Actions
- (Optional) [h2o/h2o](https://github.com/h2o/h2o)
