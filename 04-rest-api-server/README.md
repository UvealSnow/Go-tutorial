# REST API server with Go and Gin

In this tutorial, we're going to build a REST API server that hanldes the same use case as the previous tutorial, managing CRUD operations of a `albums struct`.

In the tutorial the goal is to see how easy it is to spin up an HTTP server with Go and Gin, so there's no data persistence. It uses a slice to store the albums, which means that any time the service goes down the data will be lost.

## Extensions

As additional work, I decided to add a data persistence layer, it uses the same principles as the last tutorial, I'm using a mysql database to store the data that I spin up using Docker.

```bash
$ docker run --name some-mysql \
  -e MYSQL_ROOT_PASSWORD=secret \
  -e MYSQL_USER=go-user \
  -e MYSQL_PASSWORD=go-user-pw \
  -e MYSQL_DATABASE=recordings \
  -v ./init.sql:/docker-entrypoint-initdb.d/1.sql \
  --rm -d -p 3306:3306 mysql:9.2
# Add required environment variables to the current shell
$ source .env.example
$ go run ./web-service-gin/main.go
```
