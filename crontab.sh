#!/bin/bash

cd `dirname $0`

(
    flock -x -n 200 || exit 1

	if [ -e /tmp/goweb.sock ]; then rm /tmp/goweb.sock; fi
	./goweb

) 200>/tmp/goweb.lock
