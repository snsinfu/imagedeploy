#!/usr/bin/bash -eu

set -o pipefail

while read oldrev newrev ref; do
    branch="$(git rev-parse --symbolic --abbrev-ref "${ref}")"

    case "${branch}" in
    master | main)
        ;;
    *)
        continue
    esac

    echo "* Start deploying ${branch}:${newrev}"
    (
        workdir="$(mktemp -d)"

        cleanup() {
            echo "* Cleaning up build..."
            rm -rf "${workdir}"
        }
        trap cleanup EXIT

        echo "* Checking out..."
        git archive "${newrev}" | tar -C "${workdir}" -xf -

        echo "* Working in ${workdir}"
        cd "${workdir}"

        imagedeploy
    )
    echo "* Finished!"
done
