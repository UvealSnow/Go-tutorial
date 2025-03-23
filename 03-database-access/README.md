# Database Access

## Docker

I used docker to get a MySQL database up and running. The following command will start a MySQL database that exposes the port 3306.

```bash
$ docker run --name some-mysql \
  -e MYSQL_ROOT_PASSWORD=secret \
  -e MYSQL_USER=go-user \
  -e MYSQL_PASSWORD=go-user-pw \
  -e MYSQL_DATABASE=recordings \
  -v ./init.sql:/docker-entrypoint-initdb.d/1.sql \
  --rm -d -p 3306:3306 mysql:9.2
```