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
  # The init.sql file is located in folder 03-database-access
  # Use cp to copy the file to the current folder or modify this flag
  -v ./init.sql:/docker-entrypoint-initdb.d/1.sql \
  --rm -d -p 3306:3306 mysql:9.2
# Add required environment variables to the current shell
$ source .env.example
$ go run ./web-service-gin/main.go
```

I've also added some rudimentary validation, so the server will return a 400 status code if the request body is invalid on POST and PUT requests and if the lookup ID for an album is not numeric in `/album/:id`.

## Running the server

You'll need either go or docker to run the server. You'll also need to have a MySQL database running and to update the `.env.example` file to match your DB configuration. Once that's done, you can run the server with the following command:

```bash
$ go run ./web-service-gin/main.go
```

To check out the returned data you just need to navigate to `http://localhost:3000/albums` in your browser or use curl:

```bash
$ curl http://localhost:3000/albums
```

To check the creation method you can use the following curl command:

```bash
$ curl http://localhost:3000/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "Lateralus","artist": "TOOL","price": 49.99}'
```

