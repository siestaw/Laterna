# Laterna

## API Documentation

### Base URL
`http://your-server.com/api/v1`

---
### Authentification  
The admin token will be displayed once while starting for the first time. To generate a new one, run with the `--resetAdminToken` flag.

All endpoints require the admin token in the header of the request:
```
Authorization: <token>
```

---

### Routes

#### Controller
| Method  | Route                      | Description                            | Request Body                     |
|---------|----------------------------|----------------------------------------|----------------------------------|
| GET     | `/colors/{id}`             | Get the current color of an controller | —                                |
| PUT     | `/colors/{id}`             | Set the color of an controller         |`{ "ID": "1" }`        |

---

#### Colors
| Method  | Route                      | Description                            | Request Body                     |
|---------|----------------------------|----------------------------------------|----------------------------------|
| GET     | `/colors/{id}`             | Get the current color of an controller | —                                |
| PUT     | `/colors/{id}`             | Set the color of an controller         |`{ "color": "#FF0000" }`        |

---

### cURL examples


#### Create a new controller
```bash
$ curl -X POST http://your-server.com/api/v1/controllers \
     -H "Authorization: $TOKEN"
```

#### Delete an controller
```bash
$ curl -X DELETE localhost:8080/api/v1/controllers \
     -H "Authorization: $TOKEN" \
     -d '{"ID": 1}'
```


#### Get the current color of an controller
```bash
$ curl -X GET localhost:8080/api/v1/colors/1 \
     -H "Authorization: $TOKEN"
```
#### Set the color of an controller
```bash
$ curl -X PUT localhost:8080/api/v1/colors/1 \ 
     -H "Authorization:$TOKEN" \
     -d '{"Color": "#C2C342"}'    
```
