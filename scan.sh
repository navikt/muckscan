#!/bin/sh
trufflehog --json /data | truffletool --exclude package-lock.json
