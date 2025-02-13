/*
Copyright 2022 The deepauto-io LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package core

import (
	"context"
	"fmt"
	"sync"

	akt "github.com/chatgpt-accesstoken"
)

type accessTokenStore struct {
	db   map[string]*akt.AuthExpireResult
	lock sync.RWMutex
}

func (a *accessTokenStore) Add(ctx context.Context, email string, ak *akt.AuthExpireResult) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.db[email] = ak
	return nil
}

func (a *accessTokenStore) Delete(ctx context.Context, email string) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	delete(a.db, email)
	return nil
}

func (a *accessTokenStore) Get(ctx context.Context, email string) (*akt.AuthExpireResult, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	v, ok := a.db[email]
	if !ok {
		return nil, fmt.Errorf("ak: cannot find sk")
	}
	return v, nil
}

func (a *accessTokenStore) List(ctx context.Context) (map[string]*akt.AuthExpireResult, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.db, nil
}

func NewAccessTokenStore() akt.AccessTokenStore {
	return &accessTokenStore{db: make(map[string]*akt.AuthExpireResult)}
}
