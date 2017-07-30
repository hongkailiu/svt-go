#!/bin/bash

readonly SOURCE_FOLDER="$(dirname "$(readlink -f "${0}")")"
readonly APP_FOLDER="$(dirname "$(dirname "${SOURCE_FOLDER}")")"
readonly BUILD_DIR="${APP_FOLDER}/build"
readonly VERSION_FILE="${APP_FOLDER}/conf/version"

readonly VERSION="$(git describe --tags --always --dirty)"
printf "${VERSION}" > "${VERSION_FILE}"

cp -rf "${APP_FOLDER}/conf" "${BUILD_DIR}/"
go build -o "${BUILD_DIR}/svt" github.com/hongkailiu/svt-go