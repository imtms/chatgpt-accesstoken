package core

import akt "github.com/chatgpt-accesstoken"

type FakeopenStore struct {
	URL string
}

func NewFakeopenStore() akt.FakeopenStore {
	return &FakeopenStore{URL: ""}
}

func (db *FakeopenStore) SetURL(url string) error {
	db.URL = url
	return nil
}

func (db *FakeopenStore) GetURL() string {
	return db.URL
}

func (db *FakeopenStore) DeleteURL() error {
	db.URL = ""
	return nil
}
