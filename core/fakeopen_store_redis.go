package core

import (
	akt "github.com/chatgpt-accesstoken"
	"github.com/chatgpt-accesstoken/store/redisdb"
)

const UrlKey = "fakeopenurl"

type FakeopenStoreRedis struct {
	db *redisdb.Redis
}

func NewFakeopenStoreRedis(db *redisdb.Redis) akt.FakeopenStore {
	return &FakeopenStoreRedis{db: db}
}

func (db *FakeopenStoreRedis) SetURL(url string) error {
	err := db.db.Set(UrlKey, url, 0)
	if err != nil {
		return err
	}
	return nil
}

func (db *FakeopenStoreRedis) GetURL() string {
	data := db.db.Get(UrlKey)
	return data
}

func (db *FakeopenStoreRedis) DeleteURL() error {
	err := db.db.Del(UrlKey)
	if err != nil {
		return err
	}
	return nil
}
