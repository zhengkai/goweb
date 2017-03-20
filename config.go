package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/zhengkai/goweb/handle"
	"github.com/zhengkai/goweb/module"
	"gopkg.in/ini.v1"
)

var (
	configFile = `config.ini`
)

func init() {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	configFile = dir + `/` + configFile

	fmt.Println(`load config`, configFile)

	cfg, err := ini.Load(configFile)
	if err != nil {
		panic(err)
		return
	}

	sec := cfg.Section(`site`)

	Domain = sec.Key(`domain`).String()
	if Domain == `` {
		panic(`no domain config`)
	}

	handle.SetDomain(Domain)
	module.SetDomain(Domain)

	salt := sec.Key(`salt`).String()
	if salt == `` {
		panic(`no domain config`)
	}

	Salt = []byte(salt)
	module.Salt = Salt
}
