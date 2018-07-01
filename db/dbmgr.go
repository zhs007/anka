package db

import (
	"sync"

	"github.com/zhs007/anka/base"
)

// type dbnode struct {
// 	db		Database
// 	once	sync.Once
// }

type singleton struct {
	MapDBNode map[string]*Database
	mu        sync.Mutex
}

var instance *singleton
var once sync.Once

func getDBMgr() *singleton {
	once.Do(func() {
		instance = &singleton{MapDBNode: make(map[string]*Database)}
	})

	return instance
}

// GetDBNode -
func GetDBNode(dbname string) *Database {
	dbmgr := getDBMgr()

	if v, ok := dbmgr.MapDBNode[dbname]; ok {
		return v
	}

	dbmgr.mu.Lock()
	defer dbmgr.mu.Unlock()

	if v, ok := dbmgr.MapDBNode[dbname]; ok {
		return v
	}

	dbmgr.MapDBNode[dbname] = &Database{}

	cfg := base.GetConfig()
	dbmgr.MapDBNode[dbname].OpenDB(dbname, cfg.DbPath)

	return dbmgr.MapDBNode[dbname]
}
