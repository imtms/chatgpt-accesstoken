package core

import (
	"fmt"
	"github.com/asaskevich/govalidator"
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

func (db *FakeopenStoreRedis) Set(url string) error {
	err := db.db.Set(UrlKey, url, 0)
	if err != nil {
		return err
	}
	return nil
}

func (db *FakeopenStoreRedis) Get() (string, error) {
	data := db.db.Get(UrlKey)
	if govalidator.IsNull(data) {
		return "", fmt.Errorf("fake: url is empty")
	}
	return data, nil
}

func (db *FakeopenStoreRedis) Delete() error {
	err := db.db.Del(UrlKey)
	if err != nil {
		return err
	}
	return nil
}
