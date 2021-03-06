#!/bin/sh
set -e
set -x

# Enable use of ~/.git-credentials
git config --global credential.helper store

REPO_PATH=$1
trufflehog --regex --entropy false --rules /sensitive-regex.json --json ${REPO_PATH} | truffletool
trufflehog --regex --json ${REPO_PATH} | truffletool --exclude package-lock.json
