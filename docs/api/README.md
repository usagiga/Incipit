# Index of API

## Installation API

- [POST `/api/install`](./post_install.md) Register first administrator

## Authorize API

- POST `/api/login` Login
- POST `/api/login/refresh` Refresh access token

## Admin User API

- POST `/api/admin` Create administrator
- GET `/api/admin` Get all administrators
- PATCH `/api/admin` Update administrator
- DELETE `/api/admin` Delete administrator


## Link API

- POST `/api/link` Create link
- GET `/api/link` Get all links
- GET `/api/link/shortened?short_id=x` Get specific link by short ID
- PATCH `/api/link` Update link
- DELETE `/api/link` Delete link
