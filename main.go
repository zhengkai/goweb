package main

import (
	"net"
	"net/http"
	"os"
	"syscall"

	_ "github.com/zhengkai/goweb/handle"
	_ "github.com/zhengkai/goweb/module"
)

var (
	Abc        = `def`
	socketFile = `/tmp/goweb.sock`
)

func main() {

	go func() {
		http.ListenAndServe(`:8080`, nil)
	}()

	syscall.Umask(0000)
	l, err := net.ListenUnix(`unix`, &net.UnixAddr{socketFile, `unix`})
	if err != nil {
		panic(err)
	}
	defer os.Remove(socketFile)

	err = http.Serve(l, nil)
	if err != nil {
		panic(err)
	}
}
