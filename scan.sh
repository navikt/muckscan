#!/bin/sh
trufflehog --regex --entropy false --rules /sensitive-regex.json --json /data | truffletool
trufflehog --regex --json /data | truffletool --exclude package-lock.json
