package main

import (
	"flag"
	"path"

	"github.com/zhs007/anka/base"
)

func main() {
	var curdir string
	flag.StringVar(&curdir, "path", "./", "run path")
	flag.Parse()

	base.LoadConfig(path.Join(curdir, "./config.yaml"), curdir)

	StartServ()
}
