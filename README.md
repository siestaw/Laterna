# 🛋️ Laterna (Server)

**Laterna** is a simple and lightweight REST API for viewing and setting HEX color codes over the web, written in go.

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

Make sure that your firewall supports connections to the configured port for laterna (default: `8080`), otherwise clients won't be able to connect to the API. You can do so by using ufw on most linux distributions

```sh
sudo ufw allow 8080/tcp
sudo ufw reload
```

Of course, you'll have to change `8080` to your desired port if configured otherwise


### 3. Run

Make sure that go is installed (tested with Go 1.24.5 on Linux)

```sh
make run
# OR
go run ./cmd/server/
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

| Method | Route          | Description                   | Payload             |
|--------|----------------|-------------------------------|---------------------|
| POST   | `/controllers` | Create a new controller       |                     |
| DELETE | `/controllers` | Delete an existing controller |{ "ID": 1}           |


</details>

#### 🎨 Colors

| Method | Route          | Description                                        | Payload              |
|--------|----------------|----------------------------------------------------|----------------------|
| GET    | `/colors/`     | Get the current color of all available controllers |                      |
| GET    | `/colors/{id}` | Get the current color of a specific controller     |                      |
| PUT    | `/colors/{id}` | Set the color of a controller                      | { "Color": "#FFFFFF} |

---

### Request examples
<details> <summary>📥️ HTTPie examples</summary>


#### Create a new controller

```bash
$ http POST localhost:8080/api/v1/controllers "Authorization: $TOKEN"
``` 

#### Delete a controller

```bash
$ http DELETE localhost:8080/api/v1/controllers "Authorization: $TOKEN" ID:=1
```

#### Get every controllers colors

```bash
$ http GET localhost:8080/api/v1/colors/ "Authorization: $TOKEN"
```

#### Get a specific controllers's color

```bash
$ http GET localhost:8080/api/v1/colors/1 "Authorization: $TOKEN"
```

#### Set a controller's color

```bash
$ http PUT localhost:8080/api/v1/colors/1 "Authorization: $TOKEN" Color="#C2C342"
```


</details>


<details> <summary>📥️ cURL examples</summary>


#### Create a new controller

```bash
$ curl -X POST localhost:8080/api/v1/controllers \
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


</details>

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
