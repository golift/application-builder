#!/bin/bash -x

# Deploys a new homebrew formula file to golift/homebrew-tap.
# Requires SSH credentials in ssh-agent to work.
# Run by Travis-CI when a new release is created on GitHub.
# Do not edit this file.

source .metadata.sh

make ${BINARY}.rb

git config --global user.email "${BINARY}@auto.releaser"
git config --global user.name "${BINARY}-auto-releaser"

rm -rf homebrew_release_repo
git clone git@github.com:${HBREPO}.git homebrew_release_repo

# If a bitly token file exists, we'll use that to shorten the link (and allow download counting).
if [ -f "bitly_token" ]; then
  API=https://api-ssl.bitly.com/v4/bitlinks
  # Request payload. In single quotes with double quotes escaped. :see_no_evil:
  JSON='{\"domain\": \"bit.ly\",\"title\": \"${BINARY}.v${VERSION}-${ITERATION}.tgz\", \
    \"long_url\": \"https://codeload.github.com/${GHREPO}/tar.gz/v${VERSION}\"}'
  # Request with headers and data. Using bash -c to hide token from bash -x in travis logs.
  OUT=$(bash -c "curl -s -X POST -H 'Content-type: application/json' ${API} -H \"\$(<bitly_token)\" -d \"${JSON}\"")
  # Extract link from reply.
  LINK="$(echo ${OUT} | jq -r .link | sed 's/http:/https:/')?v=v${VERSION}"
  # Replace link in formula.
  sed "s#^  url.*\$#  url \"${LINK}\"#" ${BINARY}.rb > ${BINARY}.rb.new
  if [ "$?" = "0" ] && [ "$LINK" != "null?v=v${VERSION}" ] && [ "$LINK" != "?v=v${VERSION}" ]; then
    mv ${BINARY}.rb.new ${BINARY}.rb
  fi
fi

cp ${BINARY}.rb homebrew_release_repo/Formula
pushd homebrew_release_repo
git add Formula/${BINARY}.rb
git commit -m "Update ${BINARY} on Release: v${VERSION}-${ITERATION}"
git push
popd
