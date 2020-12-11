#!/bin/bash

set -e

test -z "$GITHUB_TOKEN" && echo "GITHUB_TOKEN is not set" && exit 1
test -z "$BASE_BRANCH" && echo "BASE_BRANCH is not set" && exit 1
test -z "$HEAD_BRANCH" && echo "HEAD_BRANCH is not set" && exit 1
test -z "$MESSAGE" && echo "MESSAGE is not set" && exit 1

BASE="$BASE_BRANCH"
HEAD="$HEAD_BRANCH"

if [ -n "$BASE_OWNER" ]; then
  BASE="${BASE_OWNER}:${BASE_BRANCH}"
fi

if [ -n "$HEAD_OWNER" ]; then
  HEAD="${HEAD_OWNER}:${HEAD_BRANCH}"
fi

cd pull-request-input

wget https://github.com/github/hub/releases/download/v2.14.2/hub-linux-amd64-2.14.2.tgz -O /tmp/hub.tgz
tar zxf /tmp/hub.tgz --strip-components=2 -C /usr/local/bin hub-linux-amd64-2.14.2/bin/hub

PR=$(hub pull-request -b "$BASE" -h "$HEAD" -m "$MESSAGE")

if [ -n "$PR" ]; then
  echo "$PR" > ../pull-request-output/url.txt
fi

cat ../pull-request-output/url.txt
