# Multi-module workspaces in Go

In this tutorial, we created two modules in a multi-module workspace. This has the advantage of being able to make quick changes to the source code of one of the modules and test it in the context of the other module.

In this case, we started by creating a module with `go mod init` and adding a dependency by using the `go get` command for the `golang.org/x/example/hello/reverse` package.

The package contained functions to reverse strings, but we wanted an additional method that would reverse an integer, so we downloaded the source code of the package as as a git sub-module and modified its source code.

For Go to know about the local module, we need to run the `go work use` command that will add the path of the local module to the generated `go.work` file.
