SHELL:=/bin/bash

default:
	if [ -e /tmp/goweb.sock ]; then rm /tmp/goweb.sock; fi
	go build -o goweb main.go config.go
	./goweb

clean:
	go clean
