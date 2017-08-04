#!/bin/bash

readonly SOURCE_FOLDER="$(dirname "$(readlink -f "${0}")")"
readonly APP_FOLDER="$(dirname "$(dirname "${SOURCE_FOLDER}")")"
readonly BUILD_DIR="${APP_FOLDER}/build"
readonly BUILD_FILE="${APP_FOLDER}/build/svt"
readonly PKG_DIR_NAME="pkg"
readonly PKG_DIR="${APP_FOLDER}/build/${PKG_DIR_NAME}"

if [[ ! -f "${BUILD_FILE}" ]]; then
  echo "build_file does not exits: ${BUILD_FILE}"
  exit 1
fi

readonly VERSION=$(git describe --tags --always --dirty)
readonly GO_OS="$(uname -s)"
readonly GO_ARCH="$(uname -m)"

mkdir -p "${PKG_DIR}"
readonly PKG_BASENAME="svt-${VERSION}-${GO_OS}-${GO_ARCH}.tar.gz"
readonly PKG_FULLNAME="${BUILD_DIR}/${PKG_BASENAME}"

cp -f "${BUILD_FILE}" "${PKG_DIR}/"
cp -rf "${APP_FOLDER}/conf" "${PKG_DIR}/"
cp -rf "${APP_FOLDER}/content" "${PKG_DIR}/"

readonly CURRENT_DIR="$(pwd)"
cd "${BUILD_DIR}" || exit 1
tar -czf "${PKG_FULLNAME}" --transform "s/${PKG_DIR_NAME}/svt/" "${PKG_DIR_NAME}"
cd "${CURRENT_DIR}" || exit 1
