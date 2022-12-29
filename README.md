<h1 align="center">
  Policeman
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
  timeout: 2
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

## Endpoints

## Docker

Use **docker-compose** to run everything on docker.

```shell
docker compose up -d
```
