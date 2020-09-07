
# POST `/api/install`

Register first administrator to use Incipit.
This API cannot use not for first time.


## Request

### Header

| Key | Description |
| --- | --- |
| `Content-Type` | [Required] `application/json` |


### Body

| Key | Description |
| --- | --- |
| `name` | User name. 3-32 chars |
| `screen_name` | Screen name. 3-32 chars |
| `password` | Password. 8-72 chars |


### Example

```sh
$ curl -X POST https://example.com/api/install -H 'Content-Type:application/json' -d '{"name":"test", "screen_name": "Test Man", "password": "abcd1234"}'
```


## Response

### Body

This API returns response body as JSON.

| Key | Description |
| --- | --- |
| `type` | `install` or `error` |
| `details` | Not used in this API |
| `access_token` | Access token used in `Authorization` header |
| `refresh_token` | Refresh token used in refreshing access token |


### Example

```json
{
  "type": "install",
  "details": null,
  "access_token": "xxxxx",
  "refresh_token": "yyyyy"
}
```
