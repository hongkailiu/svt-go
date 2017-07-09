#!/bin/bash

set -e

readonly SOURCE_FOLDER="$(dirname "$(readlink -f "${0}")")"
readonly APP_FOLDER="$(dirname "$(dirname "${SOURCE_FOLDER}")")"
readonly BUILD_DIR="${APP_FOLDER}/build"
readonly RELEASE_DIR="${BUILD_DIR}/release"

readonly VERSION=$(git describe --tags --always --dirty)
readonly GO_OS="$(uname -s)"
readonly GO_ARCH="$(uname -m)"

readonly PKG_BASENAME="svt-${VERSION}-${GO_OS}-${GO_ARCH}.tar.gz"
readonly PKG_FULLNAME="${BUILD_DIR}/${PKG_BASENAME}"
if [[ ! -f "${PKG_FULLNAME}" ]]; then
  echo "pkg file does not exits: ${PKG_FULLNAME}"
  exit 1
fi

rm -rf "${RELEASE_DIR}"
mkdir -p "${RELEASE_DIR}"
readonly REPO_NAME="svt-release"
readonly GH_TOKEN="df14ce4d4468fcd4050e56e28c7ac118a8c0a6c2"
readonly REPO_URL="https://${GH_TOKEN}@github.com/cduser/${REPO_NAME}.git"

readonly CURRENT_DIR="$(pwd)"
cd "${RELEASE_DIR}" || exit 1

git clone "${REPO_URL}"
cd "${REPO_NAME}"
git checkout -b tempB
cp -f "${PKG_FULLNAME}" .
git add "${PKG_BASENAME}"
if [[ -n "${TRAVIS}" ]]; then
  echo "release by travis ci to branch: dev_${HOSTNAME}_${USERNAME}"
  msg_body="TRAVIS_BUILD_NUMBER: ${TRAVIS_BUILD_NUMBER}\nTRAVIS_BUILD_ID: ${TRAVIS_BUILD_ID}"
  git commit -m "travis: ${PKG_BASENAME}" -m "${msg_body}"
  git checkout -b
  git push origin "HEAD:travis_${TRAVIS_BUILD_NUMBER}"
else
  echo "release by dev to branch: dev_${HOSTNAME}_${USERNAME}"
  git commit -m "dev: ${PKG_BASENAME}"
  git push origin "HEAD:dev_${HOSTNAME}_${USERNAME}"
fi
cd "${CURRENT_DIR}" || exit 1