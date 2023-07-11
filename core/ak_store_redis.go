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
	"encoding/json"
	"fmt"
	akt "github.com/chatgpt-accesstoken"
	"github.com/chatgpt-accesstoken/store/redisdb"
)

const akPrefix = "ak:"

type accessTokenStoreRedis struct {
	db *redisdb.Redis
}

func (a *accessTokenStoreRedis) Add(ctx context.Context, email string, ak *akt.AuthExpireResult) error {
	data, err := json.Marshal(ak)
	if err != nil {
		return err
	}

	err = a.db.Set(akPrefix+email, string(data), 0)
	if err != nil {
		return err
	}

	return nil
}

func (a *accessTokenStoreRedis) Delete(ctx context.Context, email string) error {
	err := a.db.Del(akPrefix + email)
	if err != nil {
		return err
	}

	return nil
}

func (a *accessTokenStoreRedis) Get(ctx context.Context, email string) (*akt.AuthExpireResult, error) {

	data := a.db.Get(akPrefix + email)
	if data == "" {
		return nil, fmt.Errorf("ak: cannot find sk")
	}

	var ak akt.AuthExpireResult
	err := json.Unmarshal([]byte(data), &ak)
	if err != nil {
		return nil, err
	}

	return &ak, nil
}

func NewAccessTokenStoreRedis(db *redisdb.Redis) akt.AccessTokenStore {
	return &accessTokenStoreRedis{db: db}
}
