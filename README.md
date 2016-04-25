# replicaSetStatus

This repository contains a demo of sharing Go code with Javascript using [GopherJS](https://github.com/gopherjs/gopherjs)

[export.go](https://rawgit.com/denniskuczynski/replicaSetStatus/master/src/export.go) exposes an interface to some utility functions in the [replicaSetStatus](https://rawgit.com/denniskuczynski/replicaSetStatus/master/src/replicaSetStatus) module.

To build: `./bin/gopherjs -v build src/export.go`

To see the demo in action:
* https://rawgit.com/denniskuczynski/replicaSetStatus/master/index.html
