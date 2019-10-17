# How to learn Golang

Start with these.

1. [A tour of Golang](https://tour.golang.org/)
2. [Golang by Example](https://gobyexample.com)

Use only std packages when you are a beginner.

## Install the latest version 

1. [Install Version Manager](https://github.com/moovweb/gvm)
2. Use **$gvm install go1.13 &&gvm use go1.13 [--default]** while you refer to [the release history](https://golang.org/doc/devel/release.html).

- **$go mod init** creates a new module, initializing the go.mod file that describes it.

- $go build, $go test, and other package-building commands add new dependencies to go.mod as needed.
- **$go list -m all** prints the dependencies.
- **$go get** changes the required version of a dependency (or adds a new dependency).
- **$go mod tidy** removes unused dependencies.

## Learn how to structure the projects

1. [Use Golang moudles.](https://blog.golang.org/using-go-modules)
2. [Refer to this example.](https://github.com/golang-standards/project-layout)

## Write code

1. [Official Web Project](https://golang.org/doc/articles/wiki/), [its code](https://golang.org/doc/articles/wiki/final.go) and [this blog](http://polyglot.ninja/golang-making-http-requests/)
2. [JSON](https://www.google.com/search?&q=how+to+use+json+in+golang) and [this blog](http://polyglot.ninja/golang-json/)
3. [SQL](https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/)
4. [gRPC](https://grpc.io/docs/quickstart/go/) and make it work with module.

## READ MORE

1. [Golang Blog](https://blog.golang.org/)
