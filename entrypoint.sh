#!/bin/sh

CMD=$1
which $CMD >/dev/null 2>&1
if [ $? -eq 0 ]; then
    shift
    exec $CMD "$@"
else
    /scan.sh "$@"
fi
