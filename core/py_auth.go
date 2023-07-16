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

//
//import (
//	"context"
//	"fmt"
//	http "github.com/bogdanfinn/fhttp"
//	tls_client "github.com/bogdanfinn/tls-client"
//	"github.com/tidwall/gjson"
//	"io"
//	"time"
//)
//
//type AuthO struct {
//	email     string
//	password  string
//	proxy     string
//	session   tls_client.HttpClient
//	userAgent string
//}
//
//func NewAuth0(email, password, proxy string) *AuthO {
//	auth := &AuthO{
//		email:     email,
//		password:  password,
//		proxy:     proxy,
//		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
//	}
//
//	jar := tls_client.NewCookieJar()
//	options := []tls_client.HttpClientOption{
//		tls_client.WithTimeoutSeconds(20),
//		tls_client.WithClientProfile(tls_client.Firefox_102),
//		tls_client.WithNotFollowRedirects(),
//		tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
//		// Proxy
//		tls_client.WithProxyUrl(proxy),
//	}
//	auth.session, _ = tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
//	return auth
//}
//
//func defaultAPIPrefix() string {
//	date := time.Now().AddDate(0, 0, -1).Format("20060102")
//	return fmt.Sprintf("https://ai-%s.fakeopen.com", date)
//}
//
//func (a *AuthO) PartOne(ctx context.Context) (string, error) {
//	url := defaultAPIPrefix() + "/auth/preauth"
//
//	req, _ := http.NewRequest("GET", url, nil)
//
//	resp, err := a.session.Do(req)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		return "", fmt.Errorf("%s", string(body))
//	}
//
//	cok := gjson.Get(string(body), "preauth_cookie").String()
//	if cok == "" {
//		return "", fmt.Errorf("preauth cookie failed")
//	}
//	return cok, nil
//}
//
//func (a *AuthO) PartTwo(ctx context.Context, preauth string) (string, error) {
//	code_challenge := "w6n3Ix420Xhhu-Q5-mOOEyuPZmAsJHUbBpO8Ub7xBCY"
//	code_verifier := "yGrXROHx_VazA0uovsxKfE263LMFcrSrdm4SlC-rob8"
//
//	url := fmt.Sprintf("https://auth0.openai.com/authorize?client_id=pdlLIX2Y72MIl2rhLhTE9VV9bN905kBh&audience=https%%3A%%2F%%2Fapi.openai.com%%2Fv1&redirect_uri=com.openai.chat%%3A%%2F%%2Fauth0.openai.com%%2Fios%%2Fcom.openai.chat%%2Fcallback&scope=openid%%20email%%20profile%%20offline_access%%20model.request%%20model.read%%20organization.read%%20offline&response_type=code&code_challenge=%s&code_challenge_method=S256&prompt=login&preauth_cookie=%s", code_challenge, preauth)
//	return a.PartThree(ctx, code_verifier, url)
//}
//
//func (a *AuthO) PartThree(ctx context.Context, verifier, url string) (string, error) {
//	req, _ := http.NewRequest("GET", url, nil)
//	req.Header.Add("User-Agent", a.userAgent)
//	req.Header.Add("Referer", "https://ios.chat.openai.com/")
//
//	resp, err := a.session.Do(req)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		return "", fmt.Errorf("%s", string(body))
//	}
//
//	parsedUrl, _ := url.Parse(resp.Request.URL.String())
//	values, _ := url.ParseQuery(parsedUrl.RawQuery)
//
//}
