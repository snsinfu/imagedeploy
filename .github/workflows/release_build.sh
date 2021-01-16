#!/bin/sh -eux

tag="$(git rev-parse --symbolic --abbrev-ref "${GITHUB_REF}")"
version="${tag#v}"
os="${GOOS}"
arch="${GOARCH}"
archive="imagedeploy-${version}-${os}-${arch}.tar.gz"

go build -o imagedeploy ./main
tar -cz --numeric-owner --owner 0 --group 0 -f "${archive}" imagedeploy

echo "::set-output name=filename::${archive}"
