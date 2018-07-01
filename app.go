package main

import (
	"sync"

	"github.com/zhs007/anka/base"
	"github.com/zhs007/anka/grpcserv"
	"github.com/zhs007/anka/model/node"
)

var wg sync.WaitGroup

// StartServ -
func StartServ() {
	base.InitLogger()
	base.Info("start...")
	cfg := base.GetConfig()

	node.InitNodeMgr()
	// model.LoadGameStatistics()

	// fmt.Println(cfg.WebServAddr)
	// fmt.Println(cfg.GrpcServAddr)

	wg.Add(1)
	go grpcserv.StartServ(cfg.GrpcServAddr, &wg)

	// wg.Add(1)
	// go webserv.StartServ(cfg.WebServAddr, &wg)

	wg.Wait()
}
