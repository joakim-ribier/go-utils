# go-utils

[![Go Report Card](https://goreportcard.com/badge/github.com/joakim-ribier/go-utils)](https://goreportcard.com/report/github.com/joakim-ribier/go-utils)
![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)
[![Go Reference](https://pkg.go.dev/badge/image)](https://pkg.go.dev/github.com/joakim-ribier/go-utils)
[![codecov](https://codecov.io/gh/joakim-ribier/go-utils/graph/badge.svg?token=AUAOC8992T)](https://codecov.io/gh/joakim-ribier/go-utils)

Go library that provides utilities functions to manipulate slices type, json serialization, io, and http calls.

## Usage

The main goal is to make code cooler to write in implementations, have fun coding and discover a new language!

Run `$ go get github.com/joakim-ribier/go-utils` to add this library on your project.

Otherwise, you can download it and start `$ go build ./...` to work on.

### slicesutil

```go
// []User --> []User
slicesutil.FilterT[User](users, func(u User) bool {
	return u.Age > 17
})

// []User --> *User
slicesutil.FindT[User](users, func(u User) bool {
	return u.Age == 2
})

// []User --> []User
slicesutil.SortT[User, int](_1, func(a, b User) (int, int) {
	return a.Age, b.Age
})

// []User --> []string
slicesutil.TransformT[User, string](users, func(u User) (*string, error) {
	return &u.Name, nil
})
````

Find more examples on the [README.md (pkg/slicesutil)](pkg/slicesutil/README.md).

## Test and Coverage

```go
$ go test ./... -race -covermode=atomic -coverprofile=coverage.out
...
ok  	github.com/joakim-ribier/go-utils/pkg/cryptosutil	1.480s	coverage: 80.0% of statements
ok  	github.com/joakim-ribier/go-utils/pkg/genericsutil	1.396s	coverage: 100.0% of statements
ok  	github.com/joakim-ribier/go-utils/pkg/httpsutil	2.066s	coverage: 37.5% of statements
ok  	github.com/joakim-ribier/go-utils/pkg/iosutil	1.866s	coverage: 100.0% of statements
ok  	github.com/joakim-ribier/go-utils/pkg/jsonsutil	1.965s	coverage: 100.0% of statements
ok  	github.com/joakim-ribier/go-utils/pkg/slicesutil	1.783s	coverage: 100.0% of statements
ok  	github.com/joakim-ribier/go-utils/pkg/stringsutil	1.690s	coverage: 100.0% of statements
ok  	github.com/joakim-ribier/go-utils/pkg/timesutil	2.370s	coverage: 100.0% of statements
```

## Projects using `go-utils`

* [CLI for Postman](https://github.com/joakim-ribier/gcli-4postman)

## License
This software is licensed under the MIT license, see [License](https://github.com/joakim-ribier/go-utils/blob/main/LICENSE) for more information.