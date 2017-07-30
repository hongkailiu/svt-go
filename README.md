# svt-go

[![Build Status](https://travis-ci.org/hongkailiu/svt-go.svg?branch=master)](https://travis-ci.org/hongkailiu/svt-go)
[![Coverage Status](https://coveralls.io/repos/github/hongkailiu/svt-go/badge.svg?branch=master)](https://coveralls.io/github/hongkailiu/svt-go?branch=master)

# Prerequisites

* [Go-lang](https://golang.org/)

    ```sh
    $ go version
    go version go1.8.3 linux/amd64
    ```

* [godep](https://github.com/tools/godep)

    ```sh
    $ go get github.com/tools/godep
    ```

# Get Code

```sh
$ go get github.com/hongkailiu/svt-go
```

# Get Dependencies

```sh
$ godep restore
```

# Build and Run

```sh
$ make build
$ ./build/svt
```

# Run Tests

## run all tests

```sh
$ make test
```

## run tests in a pkg
Eg, run tests in <code>http</code> package

```sh
$ go test "github.com/hongkailiu/svt-go/http"
```

# Package

```sh
$ make clean package
$ ls build/svt*.tar.gz
```

# Release

See <code>.travis.yml</code> for details.

The packaged artifact is released to [svt-release](https://github.com/cduser/svt-release) repo.
Note that in order to activate the release we need to turn on
<code>${RELEASE}</code> on travis-ci.

