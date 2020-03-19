code-generator-deps:
	bazel run //:gazelle -- update-repos -from_file=third_party/k8s.io/code-generator/go.mod -to_macro=go/code-generator-deps.bzl%go_repositories

.PHONY: code-generator-deps