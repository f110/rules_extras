#!/usr/bin/env bash
set -e

NAME="code-generator"
VERSION="v0.17.4"

TARGETDIR="$(pwd)/third_party/${NAME}-${VERSION}"

if [ -d "${TARGETDIR}" ]; then
  rm -rf "${TARGETDIR}"
fi

TMPDIR=$(mktemp -d)
cd "${TMPDIR}"
echo "Clone: https://github.com/kubernetes/${NAME}.git"
git clone --quiet --depth 1 --branch "$VERSION" "https://github.com/kubernetes/${NAME}.git"

find "${NAME}" -name "*_test.go" -delete
find "${NAME}" -name "testdata" -type d | xargs rm -rf
find "${NAME}" -name "_examples" -type d | xargs rm -rf
cd "${NAME}"

find . -name ".*" -maxdepth 1 | grep -v "^.$" | xargs rm -rf {} +
cat <<EOS > BUILD.bazel
load("//go:vendor.bzl", "go_vendor")

# gazelle:prefix k8s.io/${NAME}

go_vendor(name = "vendor")
EOS

mv "${TMPDIR}/${NAME}" "${TARGETDIR}"
bazel run //third_party/${NAME}-${VERSION}:vendor