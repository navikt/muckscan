#!/bin/sh
set -e
#set -x

CMD=$1
if [ "$CMD" = "bash" -o "$CMD" = "sh" ];
then
    shift
    exec /bin/$CMD "$@"
else
    /scan.sh $@
fi
