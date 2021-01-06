#!/bin/sh
# Ensure this script fails if anything errors
set -e

test -z "$CHART_PATH" && (echo "CHART_PATH is not defined" && return)
test -z "$CHART_REPO" && (echo "CHART_REPO is not defined" && return)
test -z "$VERSION_FILE" && (echo "VERSION_FILE is not defined" && return)
test ! -d "$CHART_PATH" && (echo "CHART_PATH is not found" && return)

VERSION=$(cat "$VERSION_FILE")
CHART_NAME=$(basename "$CHART_PATH")
CHART_TMP_PATH="/tmp/$CHART_NAME"

apk add curl

echo "Creating package"
mkdir -p "$CHART_TMP_PATH"
cp -r "${CHART_PATH}/." "$CHART_TMP_PATH"

helm init --client-only --stable-repo-url https://charts.helm.sh/stable
helm dependency update "${CHART_TMP_PATH}"
helm package "${CHART_TMP_PATH}" --version="${VERSION}" -d /tmp

# Push chart to Artifactory
if [ -n "${ARTIFACTORY_USERNAME}" ] && [ -n "${ARTIFACTORY_PASSWORD}" ]; then
  user_args=("--user" "${ARTIFACTORY_USERNAME}:${ARTIFACTORY_PASSWORD}")
fi
CHART_PACKAGE="${CHART_NAME}-${VERSION}.tgz"
CHART_PACKAGE_PATH="/tmp/${CHART_PACKAGE}"

# Calculate checksums
CHART_PACKAGE_SHA1=$(sha1sum "${CHART_PACKAGE_PATH}" | cut -d ' ' -f 1)
CHART_PACKAGE_SHA256=$(sha256sum "${CHART_PACKAGE_PATH}" | cut -d ' ' -f 1)
CHART_PACKAGE_MD5=$(md5sum "${CHART_PACKAGE_PATH}" | cut -d ' ' -f 1)

user_args="--header X-Checksum-MD5:${CHART_PACKAGE_MD5} --header X-Checksum-Sha1:${CHART_PACKAGE_SHA1} --header X-Checksum-Sha256:${CHART_PACKAGE_SHA256}"

# Upload package
curl $user_args --fail --upload-file "${CHART_PACKAGE_PATH}" "${CHART_REPO}/${CHART_PACKAGE}"

if [ "${TAG_AS_LATEST}" = "true" ]; then
  CHART_LATEST="${CHART_NAME}-latest.tgz"
  curl $user_args --fail --upload-file "${CHART_PACKAGE_PATH}" "${CHART_REPO}/${CHART_LATEST}"
fi
