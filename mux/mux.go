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

package mux

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/chatgpt-accesstoken/render"

	akt "github.com/chatgpt-accesstoken"
	"github.com/gin-gonic/gin"
)

type Server struct {
	openAuthSvc   akt.OpenaiAuthService
	proxySvc      akt.ProxyService
	fakeopenStore akt.FakeopenStore
}

func New(openAuthSvc akt.OpenaiAuthService, proxySvc akt.ProxyService, fakeopenStore akt.FakeopenStore) *Server {
	return &Server{
		openAuthSvc:   openAuthSvc,
		proxySvc:      proxySvc,
		fakeopenStore: fakeopenStore,
	}
}

func (s Server) Handler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Any("/health", s.Healthy)

	ag := r.Group("/auth")
	{
		ag.POST("/", s.handlerPostAccessToken) // support [潘多拉]
		ag.POST("/puid", s.handlerPostPUID)
		ag.POST("/all", s.handlerPostAll)
	}

	pg := r.Group("/proxy")
	{
		pg.GET("/", s.handlerGetProxy)
		pg.POST("/", s.handlerPostProxy)
		pg.DELETE("/:id", s.handlerDeleteProxy)
	}

	pfg := r.Group("/fakeopen")
	{
		pfg.GET("/", s.handlerGetFakeopen)
		pfg.POST("/", s.hanlderPostFakeopen)
		pfg.DELETE("/", s.handlerDeleteFakeopen)
		pfg.PUT("/", s.hanlderPutFakeopen)
	}
	return r
}

func (s Server) Healthy(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (s Server) handlerPostAccessToken(ctx *gin.Context) {
	in := new(akt.OpenaiAuthRequest)
	if err := ctx.BindJSON(in); err != nil {
		render.BadRequest(ctx.Writer, err)
		return
	}

	if govalidator.IsNull(in.Email) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot find email"))
		return
	}

	if govalidator.IsNull(in.Password) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot find password"))
		return
	}

	res, err := s.openAuthSvc.AccessToken(ctx, in)
	if err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}

	render.JSON(ctx.Writer, struct {
		Default string `json:"default"`
	}{
		Default: res.AccessToken,
	}, http.StatusOK)
}

func (s Server) handlerPostPUID(ctx *gin.Context) {
	in := new(akt.OpenaiAuthRequest)
	if err := ctx.BindJSON(in); err != nil {
		render.BadRequest(ctx.Writer, err)
		return
	}

	if govalidator.IsNull(in.AccessToken) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot access token"))
		return
	}

	res, err := s.openAuthSvc.PUID(ctx, in)
	if err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}
	render.JSON(ctx.Writer, res, http.StatusOK)
}

func (s Server) handlerPostAll(ctx *gin.Context) {
	in := new(akt.OpenaiAuthRequest)
	if err := ctx.BindJSON(in); err != nil {
		render.BadRequest(ctx.Writer, err)
		return
	}

	if govalidator.IsNull(in.Email) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot find email"))
		return
	}

	if govalidator.IsNull(in.Password) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot find password"))
		return
	}

	res, err := s.openAuthSvc.AccessToken(ctx, in)
	if err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}
	render.JSON(ctx.Writer, res, http.StatusOK)
}
