# 🛋️ Laterna (Server)

## 🏵️ Setup

Laterna requires a config.json in it's root directory. An [example config](https://github.com/siestaw/Laterna/blob/main/config.json.example) is provided

`$ mv config.json.example config.json`

You can leave the json as it is or configure it to your liking, although some configuration options (e.g. `verboseLogging`) aren't fully implemented yet.

## 🛜 API Documentation

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

| Method | Route          | Description                   | Request Body  |
| ------ | -------------- | ----------------------------- | ------------- |
| POST   | `/controllers` | Create a new controller       | —             |
| DELETE | `/controllers` | Delete an existing controller | `{ "ID": 1 }` |

---

#### Colors

| Method | Route          | Description                           | Request Body             |
| ------ | -------------- | ------------------------------------- | ------------------------ |
| GET    | `/colors/{id}` | Get the current color of a controller | —                        |
| PUT    | `/colors/{id}` | Set the color of a controller         | `{ "color": "#FF0000" }` |

---

### cURL examples

#### Create a new controller

```bash
$ curl -X POST http://your-server.com/api/v1/controllers \
       -H "Authorization: $TOKEN"
```

ℹ️ Response with the newly assigned ID. The ID will always be the next available one

#### Delete a controller

```bash
$ curl -X DELETE localhost:8080/api/v1/controllers \
       -H "Authorization: $TOKEN" \
       -d '{"ID": 1}'
```

#### Get the current color of a controller

```bash
$ curl -X GET localhost:8080/api/v1/colors/1 \
       -H "Authorization: $TOKEN"
```

#### Set the color of a controller

```bash
$ curl -X PUT localhost:8080/api/v1/colors/1 \
       -H "Authorization:$TOKEN" \
       -d '{"Color": "#C2C342"}'
```
