# REST API server with Go and Gin

In this tutorial, we're going to build a REST API server that hanldes the same use case as the previous tutorial, managing CRUD operations of a `albums struct`.

In the tutorial the goal is to see how easy it is to spin up an HTTP server with Go and Gin, so there's no data persistence. It uses a slice to store the albums, which means that any time the service goes down the data will be lost.
