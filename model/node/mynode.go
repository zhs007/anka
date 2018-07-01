package node

import (
	"encoding/json"

	"github.com/zhs007/anka/db"
)

const (
	// MYNODE -
	MYNODE = "mynode"
)

var myNode Node

func loadMyNode(coredb *db.Database) (err error) {
	val, err := coredb.Get(MYNODE)
	if err != nil {
		return err
	}

	json.Unmarshal(val, &myNode)

	return nil
}

func saveMyNode(coredb *db.Database) error {
	val, err := json.Marshal(myNode)
	if err != nil {
		return err
	}

	err1 := coredb.Put(MYNODE, val)
	return err1
}
