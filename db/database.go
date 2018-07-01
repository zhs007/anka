package db

import (
	"path"

	"github.com/tecbot/gorocksdb"

	"github.com/zhs007/anka/base"
)

// Database - recoksdb
type Database struct {
	dbname string
	db     *gorocksdb.DB
	rops   *gorocksdb.ReadOptions
	wops   *gorocksdb.WriteOptions
	fops   *gorocksdb.ReadOptions
}

// FuncForeach - foreach
type FuncForeach func(key string, val []byte)

// FuncForeachString - foreach
type FuncForeachString func(key string, val string)

// FuncWriteBatch - write batch
type FuncWriteBatch func(wb *gorocksdb.WriteBatch)

// OpenDB - open database
func (db *Database) OpenDB(dbname string, dir string) error {
	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(gorocksdb.NewLRUCache(3 << 30))
	opts := gorocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	rdb, err := gorocksdb.OpenDb(opts, path.Join(dir, dbname))
	if err != nil {
		base.Error("gorocksdb.OpenDb fail")
		return err
	}
	db.db = rdb
	db.rops = gorocksdb.NewDefaultReadOptions()
	db.wops = gorocksdb.NewDefaultWriteOptions()
	db.fops = gorocksdb.NewDefaultReadOptions()
	db.fops.SetFillCache(false)

	return nil
}

// PutString - put val in key
func (db *Database) PutString(key string, val string) (err error) {
	err = db.db.Put(db.wops, []byte(key), []byte(val))

	return
}

// Put - put val in key
func (db *Database) Put(key string, val []byte) (err error) {
	err = db.db.Put(db.wops, []byte(key), val)

	return
}

// GetString - get val with key
func (db *Database) GetString(key string) (val string, err error) {
	v, e := db.db.Get(db.rops, []byte(key))
	if e != nil {
		return "", e
	}
	defer v.Free()

	val = string(v.Data())

	return val, nil
}

// Get - get val with key
func (db *Database) Get(key string) (val []byte, err error) {
	v, e := db.db.Get(db.rops, []byte(key))
	if e != nil {
		return nil, e
	}
	defer v.Free()

	// val = string(v.Data())

	return v.Data(), nil
}

// Delete - put val in key
func (db *Database) Delete(key string) (err error) {
	err = db.db.Delete(db.wops, []byte(key))

	return
}

// Foreach - put val in key
func (db *Database) Foreach(prefix string, foreach FuncForeach) (err error) {
	// ro := gorocksdb.NewDefaultReadOptions()
	// ro.SetFillCache(false)
	p := []byte(prefix)
	if db.db == nil {
		base.Info("Foreach nil")
	}
	it := db.db.NewIterator(db.fops)
	defer it.Close()

	it.Seek(p)
	for it = it; it.Valid() && it.ValidForPrefix(p); it.Next() {
		key := it.Key()
		val := it.Value()

		foreach(string(key.Data()), val.Data())

		key.Free()
		val.Free()
	}

	if err = it.Err(); err != nil {
		return
	}

	return
}

// ForeachString - put val in key
func (db *Database) ForeachString(prefix string, foreach FuncForeachString) (err error) {
	// ro := gorocksdb.NewDefaultReadOptions()
	// ro.SetFillCache(false)
	p := []byte(prefix)
	if db.db == nil {
		base.Info("Foreach nil")
	}
	it := db.db.NewIterator(db.fops)
	defer it.Close()

	it.Seek(p)
	for it = it; it.Valid() && it.ValidForPrefix(p); it.Next() {
		key := it.Key()
		val := it.Value()

		foreach(string(key.Data()), string(val.Data()))

		key.Free()
		val.Free()
	}

	if err = it.Err(); err != nil {
		return
	}

	return
}

// WriteBatch - WriteBatch
func (db *Database) WriteBatch(writebatch FuncWriteBatch) (err error) {
	wb := gorocksdb.NewWriteBatch()
	defer wb.Destroy()
	writebatch(wb)
	err = db.db.Write(db.wops, wb)
	return
}
