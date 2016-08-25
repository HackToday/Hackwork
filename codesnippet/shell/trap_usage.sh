#!  /usr/bin/env bash

trap "echo Boo" SIGINT SIGTERM

echo "pid is $$"

while :
do
	sleep 2
done
