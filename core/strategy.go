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
	"errors"
	"fmt"
	"github.com/chatgpt-accesstoken/store/redisdb"
	"math/rand"
	"sync"
	"time"
)

type StrategyBalance interface {
	// Select obtain IP according to selection strategy
	Select(service []string) (string, error)
}

type RandomStrategy struct{}

func (s *RandomStrategy) Select(ips []string) (string, error) {
	if ips == nil || len(ips) == 0 {
		return "", errors.New("strategy: ip not exist")
	}

	return ips[rand.Intn(len(ips))], nil
}

type localExpireStrategy struct {
	proxysMap map[string]int64
	expire    time.Duration // 5秒或10秒
	lock      sync.RWMutex
}

func NewLocalExpireStrategy(expire time.Duration) StrategyBalance {
	return &localExpireStrategy{
		proxysMap: make(map[string]int64),
		expire:    expire,
	}
}

func (s *localExpireStrategy) Select(ips []string) (string, error) {
	currentTime := time.Now().Unix()

	for _, ip := range ips {
		s.lock.Lock()
		val, ok := s.proxysMap[ip]
		if !ok || currentTime > val+int64(s.expire.Seconds()) {
			s.proxysMap[ip] = currentTime
			s.lock.Unlock()
			return ip, nil
		}
		s.lock.Unlock()
	}
	return "", fmt.Errorf("select: no available IP found")
}

type redisExpireStrategy struct {
	db     *redisdb.Redis
	expire time.Duration
}
