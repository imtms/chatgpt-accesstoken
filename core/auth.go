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
	"errors"
	"github.com/acheong08/OpenAIAuth/auth"
	akt "github.com/chatgpt-accesstoken"
	"github.com/chatgpt-accesstoken/newauth"
	"github.com/linweiyuan/go-chatgpt-api/api"
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
	var loginInfo = newauth.LoginInfo{
		Username: req.Email,
		Password: req.Password,
	}
	userLogin := newauth.UserLogin{
		Client: api.NewHttpClient(),
	}
	err := userLogin.Client.SetProxy(req.Proxy)
	if err != nil {
		return nil, err
	}
	token, err := userLogin.Login(loginInfo)
	if err != nil {
		return nil, err
	}
	responseMap := make(map[string]string)
	err = json.Unmarshal(token, &responseMap)
	if err != nil {
		return nil, err
	}
	result := auth.AuthResult{
		AccessToken:  responseMap["accessToken"],
		RefreshToken: "",
		PUID:         "",
	}
	return &result, nil
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
