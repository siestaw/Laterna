# 🛋️ Laterna (Server)

**Laterna** is a lightweight REST API for sharing HEX color codes between 2 (or more) lamps, written in go.

> [!NOTE]
> This project was developed with my own use case in mind.
> This is my first time building a REST API and my first time utilizing the go language.
> Expect bugs and weird annoyances/limitations

## ⚙️ Setup

### 1. Clone the repository

```sh
git clone https://github.com/siestaw/laterna.git
cd laterna
```

### 2. Configure the server

Laterna requires a config.json file in the root directory. A [example config](https://github.com/siestaw/Laterna/blob/main/config.json.example) is provided

```sh
cp config.json.example config.json
```

You can leave the json as it is or configure it to your liking, although some configuration options (e.g. `verboseLogging`) aren't fully implemented yet.

### 3. Run

Make sure that go is installed (tested with Go 1.24.5 on Linux)

```sh
make run
```

For advanced Makefile options:

```sh
make help
```

## 🛜 API Documentation

### 🔑 Authentication

All requests require the **admin token** via the Authorization header:

```
Authorization: <token>
```

The token is shown once on the first startup. To regenerate it, run the server with:

```sh
./laterna --resetAdminToken
```

---

### 📦️ Endpoints

#### 🎛️ Controllers

| Method | Route          | Description                   | Request Body  |
| ------ | -------------- | ----------------------------- | ------------- |
| POST   | `/controllers` | Create a new controller       | —             |
| DELETE | `/controllers` | Delete an existing controller | `{ "ID": 1 }` |

<details> <summary>🔧 Example Payload</summary>

```jsonc
// DELETE /controllers
{
    "ID": 1
}
```

</details>

#### 🎨 Colors

| Method | Route          | Description                                        | Request Body             |
| ------ | -------------- | -------------------------------------------------- | ------------------------ |
| GET    | `/colors/`     | Get the current color of all available controllers | —                        |
| GET    | `/colors/{id}` | Get the current color of a specific controller     | —                        |
| PUT    | `/colors/{id}` | Set the color of a controller                      | `{ "color": "#FF0000" }` |

<details> <summary>🔧 Example Payload</summary>

```jsonc
// PUT /colors/1
{
    "Color": "#5398B7"
}
```

</details>

---

### 📥️ cURL examples

#### Create a new controller

```bash
$ curl -X POST http://your-server.com/api/v1/controllers \
       -H "Authorization: $TOKEN"
```

#### Delete a controller

```bash
$ curl -X DELETE localhost:8080/api/v1/controllers \
       -H "Authorization: $TOKEN" \
       -d '{"ID": 1}'
```

#### Get every controllers colors

```bash
$ curl -X GET localhost:8080/api/v1/colors/ \
       -H "Authorization: $TOKEN"
```

#### Get a specific controllers's color

```bash
$ curl -X GET localhost:8080/api/v1/colors/1 \
       -H "Authorization: $TOKEN"
```

#### Set a controller's color

```bash
$ curl -X PUT localhost:8080/api/v1/colors/1 \
       -H "Authorization:$TOKEN" \
       -d '{"Color": "#C2C342"}'
```

---

### 📤️ Response Format

All responses follow this format:

```json
{
    "data": {
        "color": "#C16A31",
        "id": 1,
        "updated_at": "2025-07-28T19:47:46Z"
    },
    "status": 200,
    "success": true,
    "timestamp": "2025-07-28T20:40:58Z"
}
```

In case of an error:

```json
{
    "error": "Invalid token",
    "status": 401,
    "success": false,
    "timestamp": "2025-07-28T20:41:54Z"
}
```
