
# Error

## Response Body Format

If error occurred, API returns response body as JSON.

| Key | Description |
| --- | --- |
| `type` | `install` or `error` |
| `details` | Not used in this API |
| `p_code` | Primary error code. This represents region of error. |
| `s_code` | Secondary error code. This represents detail of error. |

To know error codes, see Error Code Index in this page.


### Example

```json
{
  "type": "error",
  "details": null,
  "p_code": 101,
  "s_code": 101
}
```


## Error Code Index

### 000 EXAMPLE

- 101 Example error

In this case, `000` is Primary Error Code.
`101` is Secondary Error Code.

### 101 Admin User Validation

- 101 Name is too short. Must be 3 or more chars.
- 102 Name is too long. Must be 32 or fewer chars.
- 103 Name has unavailable char. Must be `[a-zA-Z0-9_]+`
- 201 Screen name is too short. Must be 3 or more chars.
- 202 Screen name is too long. Must be 32 or fewer chars.
- 203 Screen name has unavailable char. Must be `[a-zA-Z0-9_]+`
- 301 Password is too short. Must be 3 or more chars.
- 302 Password is too long. Must be 32 or fewer chars.
- 303 Password has unavailable char. Must be `[a-zA-Z0-9_]+`

### 102 Link Validation

- 101 URL is Incipit. Mustn't be itself.
- 102 URL is invalid.

### 201 CRED Admin

- 101 Failed add.
- 201 Failed find.
- 202 Finding user not found.
- 301 Failed update.
- 302 Updating user not found.
- 401 Failed delete.

### 202 Authorization(Admin)

- 101 Failed to find user.
- 102 Unmatch password.
- 103 Access token is expired.
- 104 Failed to store generated token.

### 203 Hash

- 101 Failed generate hash.
- 201 Failed compare hash.

### 204 Installer

None yet.

### 205 CRED Link

- 101 Failed add.
- 201 Failed find.
- 202 Finding link not found.
- 301 Failed update.
- 302 Updating link not found.
- 401 Failed delete.

### 301 Handle request of Admin User

- 101 Failed to bind JSON. Request body is not JSON or invalid.

### 302 Handle request of Authorization(Admin)

- 101 Failed to bind JSON. Request body is not JSON or invalid.

### 303 Handle request of Link

- 101 Failed to bind JSON. Request body is not JSON or invalid.

### 304 Handle request of Install

- 101 Failed to bind JSON. Request body is not JSON or invalid.
