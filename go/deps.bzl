load("@bazel_gazelle//:deps.bzl", "go_repository")
load(":code-generator-deps.bzl", "go_repositories")

def go_extras_dependencies():
    go_repository(
        name = "io_k8s_code_generator",
        importpath = "k8s.io/code-generator",
        urls = ["https://github.com/kubernetes/code-generator/archive/v0.17.4.zip"],
        strip_prefix = "code-generator-0.17.4",
        type = "zip",
        build_file_proto_mode = "disable_global",
    )

    go_repositories()
