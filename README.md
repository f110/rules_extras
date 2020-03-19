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

### /go/gen.bzl

Vendoring generated codes by k8s.io/code-generator into the repository.

Currently supported generator is deepcopy-gen, client-gen, lister-gen and informer-gen only.

WORKSPACE:

```starlark
load("@dev_f110_rules_extras//go:deps.bzl", "go_extras_dependencies")

go_extras_dependencies()
```

```starlark
load("@dev_f110_rules_extras//go:gen.bzl", "k8s_code_generator")

k8s_code_generator(
    name = "gen",
    srcs = [
        "//pkg/api/etcd/v1alpha1:go_default_library",
    ],
    clientpackage = "github.com/f110/rules_extras/pkg/client",
    header = ":boilerplate.go.txt",
    informerpackage = "github.com/f110/rules_extras/pkg/informers",
    listerpackage = "github.com/f110/rules_extras/pkg/listers",
)
```

# Author

Fumihiro Ito