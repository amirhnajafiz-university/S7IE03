<h1 align="center">
  S7IE03
</h1>

<br />

First final project of Internet Engineering Course. 
Building a **Http Endpoint Monitoring** service with **Golang** and **Mongodb**. 
Implementing a **Restful API** to monitor users endpoints.

## Run application

The application is build with **Golang Cobra** library. In order to run the application
you need to build it first with the following command.

```shell
go build -o main
```

Output of the above command is a ```main``` file which is executable.

Execute the file with following commands:

```shell
chmod +x ./main
./main
```

### Configs

Configs will be read from ```config.yml``` file by **Golang Koanf** library.
To set your configs, copy the example file, ```cp config.example.yml config.yml```.

```yaml
# http service port
http_port: 8080
# failed requests' threshold for each endpoint
threshold: 20
# maximum endpoints that each user can create
user_endpoints: 20
# jwt authentication configs
jwt:
  private_key: "mysupersecretkey"
  expire_time: 30
# mongodb cluster configs
mongodb:
  database: "policeman"
  host: "localhost"
  port: 27017
# worker options
worker:
  timeout: 2 # seconds
  workers: 5
```

### HTTP

To run the HTTP server (which is implemented by **Golang Fiber** framework), use the
following command:

```shell
./main http
```

This command will start http server on ```localhost:[http_port]```.

### Worker

Inorder to run the worker that monitors the endpoints, use the following command:

```shell
./main worker
```

## Database

This application uses **MongoDB** for its database.

### Models

#### User

```go
type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}
```

#### Endpoint

```go
type Endpoint struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Username    string             `bson:"username"`
	Url         string             `bson:"url"`
	Threshold   int                `bson:"threshold"`
	FailedTimes int                `bson:"failedTimes"`
	CreateTime  time.Time          `json:"create_time"`
}
```

#### Request

```go
type Request struct {
	EndpointId string    `bson:"endpoint_id"`
	Code       int       `bson:"result"`
	CreateTime time.Time `bson:"create_time"`
}
```

## HTTP requests

### User

#### Register

```shell
curl --request POST \
  --url http://localhost:8080/api/register \
  --header 'Content-Type: application/json' \
  --data '{
	"username": "amirhossein",
	"password": "amirhossein"
}'
```

```shell
OK
```

#### Login

```shell
curl --request POST \
  --url http://localhost:8080/api/login \
  --header 'Content-Type: application/json' \
  --data '{
	"username": "amirhossein",
	"password": "amirhossein"
}'
```

```json
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzIyNjY5NDEsInBhc3N3b3JkIjoiJDJhJDEwJHlIQ3cueDQvU0M0N3BOTjhIY3VmS084M3A4eWpiNXhqVkJyZnYzVldDMm9hV0NmQjdVZXZ1IiwidXNlcm5hbWUiOiJhbWlyaG9zc2VpbiJ9.12pVs9ncLM4EfaHH4GkEoi44Zz4x6aqTc0T17XBFBiE",
	"expires": "2022-12-29T02:05:41.848018+03:30"
}
```

### Endpoint

#### Register

```shell
curl --request POST \
  --url http://localhost:8080/api/endpoints \
  --header 'Content-Type: application/json' \
  --header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzIyNjUwODMsInBhc3N3b3JkIjoiJDJhJDEwJHlIQ3cueDQvU0M0N3BOTjhIY3VmS084M3A4eWpiNXhqVkJyZnYzVldDMm9hV0NmQjdVZXZ1IiwidXNlcm5hbWUiOiJhbWlyaG9zc2VpbiJ9.VS3XydzeLovP7DLfVrn9B4fkDv0PnZuuzPlY2JNFOVc' \
  --data '{
	"address": "https://bale.ai"
}'
```

```shell
63acbac72a7a235b16a5e228
```

#### Get endpoints

```shell
curl --request GET \
  --url http://localhost:8080/api/endpoints \
  --header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzIyNjY5NDEsInBhc3N3b3JkIjoiJDJhJDEwJHlIQ3cueDQvU0M0N3BOTjhIY3VmS084M3A4eWpiNXhqVkJyZnYzVldDMm9hV0NmQjdVZXZ1IiwidXNlcm5hbWUiOiJhbWlyaG9zc2VpbiJ9.12pVs9ncLM4EfaHH4GkEoi44Zz4x6aqTc0T17XBFBiE'
```

```json
[
	{
		"id": "63acbac72a7a235b16a5e228",
		"address": "https://bale.ai",
		"created_at": "2022-12-28T21:53:11.199Z"
	}
]
```

#### Get requests of endpoint

```shell
curl --request GET \
  --url http://localhost:8080/api/endpoint/63acbac72a7a235b16a5e228 \
  --header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzIyNjY5NDEsInBhc3N3b3JkIjoiJDJhJDEwJHlIQ3cueDQvU0M0N3BOTjhIY3VmS084M3A4eWpiNXhqVkJyZnYzVldDMm9hV0NmQjdVZXZ1IiwidXNlcm5hbWUiOiJhbWlyaG9zc2VpbiJ9.12pVs9ncLM4EfaHH4GkEoi44Zz4x6aqTc0T17XBFBiE'
```

```json
[
	{
		"status": 504,
		"time": "2022-12-28T22:14:33.035Z"
	},
	{
		"status": 200,
		"time": "2022-12-28T22:14:33.035Z"
	}
]
```

#### Get endpoint warnings

```shell
curl --request GET \
  --url http://localhost:8080/api/endpoint/63acba882a7a235b16a5e227/warnings \
  --header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzIyNjY5NDEsInBhc3N3b3JkIjoiJDJhJDEwJHlIQ3cueDQvU0M0N3BOTjhIY3VmS084M3A4eWpiNXhqVkJyZnYzVldDMm9hV0NmQjdVZXZ1IiwidXNlcm5hbWUiOiJhbWlyaG9zc2VpbiJ9.12pVs9ncLM4EfaHH4GkEoi44Zz4x6aqTc0T17XBFBiE'
```

```json
{
	"address": "http://snapp.ir",
	"message": "this endpoint is fine."
}
```

#### Remove endpoint

```shell
curl --request DELETE \
  --url http://localhost:8080/api/endpoint/000000000000000000000000 \
  --header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzIyNjY5NDEsInBhc3N3b3JkIjoiJDJhJDEwJHlIQ3cueDQvU0M0N3BOTjhIY3VmS084M3A4eWpiNXhqVkJyZnYzVldDMm9hV0NmQjdVZXZ1IiwidXNlcm5hbWUiOiJhbWlyaG9zc2VpbiJ9.12pVs9ncLM4EfaHH4GkEoi44Zz4x6aqTc0T17XBFBiE'
```

## Docker

Use **docker-compose** to run everything on docker.

```shell
docker compose up -d
```
