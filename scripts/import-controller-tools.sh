#!/usr/bin/env bash
set -e

NAME="controller-tools"
VERSION="v0.2.4"

TARGETDIR="$(pwd)/third_party/${NAME}-${VERSION}"

if [ -d "${TARGETDIR}" ]; then
  rm -rf "${TARGETDIR}"
fi

TMPDIR=$(mktemp -d)
cd "${TMPDIR}"
echo "Clone: https://github.com/kubernetes-sigs/controller-tools.git"
git clone --depth 1 --branch "$VERSION" https://github.com/kubernetes-sigs/controller-tools.git
find "${NAME}" -name "*_test.go" -delete
find "${NAME}" -name "testdata" -type d | xargs rm -rf
cd "${NAME}"

find . -name ".*" -maxdepth 1 | grep -v "^.$" | xargs rm -rf {} +
cat <<EOS > BUILD.bazel
load("//go:vendor.bzl", "go_vendor")

# gazelle:prefix sigs.k8s.io/${NAME}

go_vendor(name = "vendor")
EOS

mv "${TMPDIR}/${NAME}" "${TARGETDIR}"
bazel run //third_party/${NAME}-${VERSION}:vendor