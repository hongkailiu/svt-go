# svt-go

[![Build Status](https://travis-ci.org/hongkailiu/svt-go.svg?branch=master)](https://travis-ci.org/hongkailiu/svt-go)
[![Coverage Status](https://coveralls.io/repos/github/hongkailiu/svt-go/badge.svg?branch=master)](https://coveralls.io/github/hongkailiu/svt-go?branch=master)

## Prerequisites

* [Go-lang](https://golang.org/): <code>${GOROOT}/bin</code> is [configure](https://golang.org/doc/install#install) as a part of <code>${PATH}</code>. 

    ```sh
    $ go version
    go version go1.8.3 linux/amd64
    ```

* [godep](https://github.com/tools/godep)

    ```sh
    $ go get github.com/tools/godep
    ```

* [ginkgo](https://onsi.github.io/ginkgo/): Needed only for test.

    ```sh
    $ go get github.com/onsi/ginkgo/ginkgo
    ```


## Get Code

```sh
$ go get github.com/hongkailiu/svt-go
```

## Get Dependencies

```sh
$ godep restore
```

## Build and Run

```sh
$ make build
$ ./build/svt
```

### Cluster Loader
For example, we can run node-virtical test by:

```sh
# ./build/svt clusterLoader --file=conf/nodeVertical.yaml
```

See [doc](doc/cluster_loader.md) for more information.


### Inventory File Generator

```sh
$ AWS_ACCESS_KEY_ID=aaa AWS_SECRET_ACCESS_KEY=bbb ./build/svt invGen
```

### Web app
A web application implemented with GoLang:

```sh
$ ./build/svt http
$ # in another terminal
$ curl localhost:8080
{
  "version": "0.0.2-3-g2375151-dirty",
  "ips": ["127.0.0.1", "::1", "192.168.31.163", "fe80::f2d5:bfff:fe5c:1b01", "192.168.122.1", "10.10.120.59"],
  "now": "2017-08-05T14:48:09.441967753-04:00"
}
```


It can log according posting http requests, to test oc-logging, and
report the host ips, to test loading-balancing of oc-services.

The docker image containing it is pushed to docker hub:

```sh
$ docker run -d -p 8080:8080 docker.io/hongkailiu/svt-go:http
```

## Run Tests

### run all tests

```sh
$ make test
```

### run tests in a pkg
Eg, run tests in <code>http</code> package

```sh
$ go test "github.com/hongkailiu/svt-go/http"
```

## Package

```sh
$ make clean package
$ ls build/svt*.tar.gz
```

## Release

See <code>.travis.yml</code> for details.

The packaged artifact is released to
[svt-release](https://github.com/cduser/svt-release) repo.
Note that in order to activate the release we need to turn on
<code>${RELEASE}</code> on travis-ci.

## Try the release version
See [the wiki page](https://github.com/hongkailiu/svt-go/wiki).

## Extended test
Functional tests on svt-go. It uses test framework _ginkgo_.
Check [Makefile](Makefile) to see how to build and run extended test.
