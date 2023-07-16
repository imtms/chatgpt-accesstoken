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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/tidwall/gjson"
	"io/ioutil"
	systemhttp "net/http"

	"github.com/acheong08/OpenAIAuth/auth"
	akt "github.com/chatgpt-accesstoken"
)

type openaiAuthService struct{}

func New() akt.OpenaiAuthService {
	return &openaiAuthService{}
}

func (s *openaiAuthService) All(ctx context.Context, req *akt.OpenaiAuthRequest) (*auth.AuthResult, error) {
	authenticator := auth.NewAuthenticator(req.Email, req.Password, req.Proxy)
	if err := authenticator.Begin(); err != nil {
		return nil, OError{Err: err}
	}

	resp := authenticator.GetAuthResult()
	if resp.AccessToken == "" || resp.PUID == "" {
		return nil, errors.New("access_token or puid is empty")
	}
	return &resp, nil
}

func (s *openaiAuthService) AccessToken(ctx context.Context, req *akt.OpenaiAuthRequest) (*auth.AuthResult, error) {
	payloadBytes, err := json.Marshal(struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Proxy    string `json:"proxy"`
	}{
		Email:    req.Email,
		Password: req.Password,
		Proxy:    req.Proxy,
	})

	body := bytes.NewReader(payloadBytes)

	kitreq, err := systemhttp.NewRequest("POST", "http://149.28.231.64:5000/api/auth", body)
	if err != nil {
		return nil, err
	}

	kitreq.Header.Set("Content-Type", "application/json")
	resp, err := systemhttp.DefaultClient.Do(kitreq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", string(responseData))
	}

	token := gjson.Get(string(responseData), "default").String()
	if govalidator.IsNull(token) {
		return nil, fmt.Errorf("token is empty")
	}

	authenticator := auth.NewAuthenticator(req.Email, req.Password, req.Proxy)
	authenticator.AuthResult.AccessToken = token
	return &authenticator.AuthResult, nil

	//authenticator := auth.NewAuthenticator(req.Email, req.Password, req.Proxy)
	//if err := authenticator.Begin(); err != nil {
	//	return nil, OError{Err: err}
	//}
	//
	//return &authenticator.AuthResult, nil
}

func (s *openaiAuthService) PUID(ctx context.Context, req *akt.OpenaiAuthRequest) (*auth.AuthResult, error) {
	authenticator := auth.NewAuthenticator(req.Email, req.Password, req.Proxy)
	authenticator.AuthResult.AccessToken = req.AccessToken

	puid, err := authenticator.GetPUID()
	if err != nil {
		return nil, OError{Err: err}
	}

	return &auth.AuthResult{
		AccessToken: req.AccessToken,
		PUID:        puid,
	}, nil
}
