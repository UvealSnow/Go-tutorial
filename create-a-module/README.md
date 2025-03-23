# Go tutorial - create a module

This code represents the result of a simple tutorial on the Go programming language. The tutorial is available at [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module).

It consists of two parts, an importable module called "greetings" and a "hello" program that uses the module.

## Requisites

- Go 1.21 or later
- A command-line interface

## Usage

To run the "hello" program, execute the following command inside of the "hello" directory:

```bash
go run .
```

## Changes

- The original tutorial uses hardcoded names. I've added stdin reading to allow the user to input a name or a list of comma-separated names.
