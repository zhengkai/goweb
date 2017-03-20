#!/bin/bash

(
    flock -x -n 200 || exit 1

	if [ -e /tmp/goweb.sock ]; then rm /tmp/goweb.sock; fi
	./goweb

) 200>/tmp/goweb.lock
