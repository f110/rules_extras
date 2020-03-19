load("@bazel_gazelle//:deps.bzl", "go_repository")
load(":code-generator-deps.bzl", "go_repositories")

def go_extras_dependencies():
    go_repository(
        name = "io_kubernetes_code_generator",
        importpath = "k8s.io/code-generator",
        urls = ["https://github.com/kubernetes/code-generator/archive/v0.17.4.zip"],
        strip_prefix = "code-generator-0.17.4",
        type = "zip",
    )

    go_repositories()
