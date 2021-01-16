#!/bin/sh -eux

version="${GITHUB_REF#v}"
os="${GOOS}"
arch="${GOARCH}"
archive="imagedeploy-${version}-${os}-${arch}.tar.gz"

go build -o imagedeploy .
tar -cz --numeric-owner --owner 0 --group 0 -f "${archive}" imagedeploy

echo "::set-output name=filename::${archive}"
