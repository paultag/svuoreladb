package main

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type SvuorelaDB struct {
	db *leveldb.DB
}

func LoadDatabase(path string) (*SvuorelaDB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &SvuorelaDB{
		db: db,
	}, nil
}

func (database *SvuorelaDB) Get(value string) *string {
	data, err := database.db.Get([]byte(value), nil)
	if err != nil {
		return nil
	}
	ret := string(data)
	return &ret
}

func (database *SvuorelaDB) Write(key, value string) {
	database.db.Put([]byte(key), []byte(value), nil)
}
