#!/bin/bash

readonly SOURCE_FOLDER="$(dirname "$(readlink -f ${0})")"
readonly APP_FOLDER=$(dirname $(dirname "${SOURCE_FOLDER}"))
readonly BUILD_DIR="${APP_FOLDER}/build"
readonly ACC_FILE="${BUILD_DIR}/coverage/acc.out"

if [[ ! -f "${ACC_FILE}" ]]; then
  echo "acc_file does not exits: ${ACC_FILE}"
fi

if [[ -n "${COVERALLS}" ]]; then
  echo "uploading results (${ACC_FILE}) to coveralls.io ..."
  cat ${ACC_FILE}
  "${HOME}/gopath/bin/goveralls" -coverprofile="${ACC_FILE}" -service travis-ci
fi
