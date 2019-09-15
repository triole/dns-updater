#!/bin/bash

# IFS=$'\n'
scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
maindir="$(echo "${scriptdir}" | grep -Po ".*(?=\/)")"
conf="${scriptdir}/targets.yaml"

SOURCE_DIR="${1}"
APP_NAME="${2}"
TARGET_FOLDER="${maindir}/${3}"
AUTHOR="${4}"
GIT_COMMIT_NO=$(git rev-list --all --count)
GIT_COMMIT_HASH=$(git rev-parse HEAD)
DATE=$(date "+%Y-%m-%d %H:%M:%S %Z")

SOURCE_DIR="${maindir}/${SOURCE_DIR}"

echo ""
while read line; do
    arch=$(echo "${line}" | grep -Po "^[a-zA-Z0-9_]+")
    flags=$(echo "${line}" | grep -Po "(?<=:).*")

    cmd="CGO_ENABLED=0 ${flags}
    	go build -ldflags \"-s -w -X 'main.BUILDTAGS={ \
    		_subversion: ${GIT_COMMIT_NO}, \
    		Author: ${AUTHOR}, \
    		Build date: ${DATE}, \
    		Git hash: ${GIT_COMMIT_HASH} \
    	}'\" \
    	-a -o \"${TARGET_FOLDER}/${arch}/${APP_NAME}\" ${SOURCE_DIR}/*.go"

    echo "Build \"${arch}/${APP_NAME}\""
    eval "${cmd}"
    if [[ -d "${maindir}/config" ]]; then
        cp -f ${maindir}/config/* "${TARGET_FOLDER}/${arch}/"
    fi
done < "${conf}"

echo ""
