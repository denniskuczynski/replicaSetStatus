# replicaSetStatus

This repository contains a demo of sharing Go code with Javascript using [GopherJS](https://github.com/gopherjs/gopherjs)

[export.go](./src/export.go) exposes an interface to some utility functions in the [replicaSetStatus](./src/replicaSetStatus) module.

To build: `./bin/gopherjs -m build src/export.go`

To see the demo for visualizing replica set sync chaining visit:
* https://rawgit.com/denniskuczynski/replicaSetStatus/master/index.html
