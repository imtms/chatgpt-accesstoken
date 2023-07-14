package core

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	akt "github.com/chatgpt-accesstoken"
)

type FakeopenStore struct {
	URL string
}

func NewFakeopenStore() akt.FakeopenStore {
	return &FakeopenStore{URL: ""}
}

func (db *FakeopenStore) Set(url string) error {
	db.URL = url
	return nil
}

func (db *FakeopenStore) Get() (string, error) {
	if govalidator.IsNull(db.URL) {
		return "", fmt.Errorf("fake: url is empty")
	}
	return db.URL, nil
}

func (db *FakeopenStore) Delete() error {
	db.URL = ""
	return nil
}
