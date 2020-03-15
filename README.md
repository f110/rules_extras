rules_extras
---

This repository contains some useful bazel's rule for my workflow.

## go

### /go/vendor.bzl

Vendoring dependent modules into the repository.

```starlark
load("@dev_f110_rules_extras//go:vendor.bzl", "go_vendor")

go_vendor(name = "vendor")
```

### /go/grpc.bzl

Vendoring generated codes by protoc-gen-grpc into the repository.

```starlark
load("@dev_f110_rules_extras//go:grpc.bzl", "vendor_grpc_source")

vendor_grpc_source(name = "vendor")
```

# Author

Fumihiro Ito