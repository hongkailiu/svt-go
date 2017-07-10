#!/bin/bash

# workaround for issue https://github.com/mattn/goveralls/issues/20
# https://raw.githubusercontent.com/gopns/gopns/master/test-coverage.sh

readonly SOURCE_FOLDER="$(dirname "$(readlink -f "${0}")")"
readonly APP_FOLDER="$(dirname "$(dirname "${SOURCE_FOLDER}")")"
readonly BUILD_DIR="${APP_FOLDER}/build"
readonly TARGET_DIR="${BUILD_DIR}/coverage"
readonly ACC_FILE="${BUILD_DIR}/coverage/acc.out"
readonly PROFILE_FILE="${BUILD_DIR}/coverage/profile.out"
readonly CURRENT_DIR="$(pwd)"

mkdir -p "${TARGET_DIR}"

echo "mode: set" > "${ACC_FILE}"

cd "${APP_FOLDER}" || exit 1
while read -r dir;
do
    if ls "${dir}"/*.go &> /dev/null; then
        return_val="$(go test -coverprofile="${PROFILE_FILE}" "${dir}")"
        echo "${return_val}"
        if [[ ${return_val} != *FAIL* ]]; then
            if [[ -f "${PROFILE_FILE}" ]]; then
                grep -v "mode: set" "${PROFILE_FILE}" >> "${ACC_FILE}"
            fi
        else
            exit 1
        fi
    fi
done << EOF
$(find . -maxdepth 10 -type d -not -path "./vendor*")
EOF
cd "${CURRENT_DIR}" || exit 1
