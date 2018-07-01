package node

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/zhs007/anka/base"
	"github.com/zhs007/anka/db"
)

const (
	// PREFIX -
	PREFIX = "node:"
)

// Node -
type Node struct {
	NameID   string `json:"nameid"`
	ServAddr string `json:"servaddr"`
	LastTime int64  `json:"lasttime"`
}

var mapNode map[string]*Node
var onceMapNode sync.Once

func loadAllNodes() (err error) {
	mapNode = make(map[string]*Node)

	cfg := base.GetConfig()
	coredb := db.GetDBNode(cfg.CoreDbName)
	if coredb == nil {
		base.Info("coredb is nil")
	}

	err = coredb.Foreach(PREFIX, func(key string, val []byte) {
		lststr := strings.Split(key, ":")
		if len(lststr) == 2 {
			nameid := lststr[1]

			mapNode[nameid] = &Node{NameID: nameid}

			json.Unmarshal(val, mapNode[nameid])
		}
	})
	if err != nil {
		return
	}

	return
}

// InitNodeMgr -
func InitNodeMgr() (err error) {
	onceMapNode.Do(func() {
		err = loadAllNodes()
	})

	return
}
