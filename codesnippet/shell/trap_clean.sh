#! /usr/bin/env bash

set -x

LOCKFILE=/var/lock/my.lock

[ -f $LOCKFILE ] && exit 0

function cleanup() {
	rm -f $LOCKFILE
	exit 255
}

trap cleanup EXIT

touch $LOCKFILE

exit 0
