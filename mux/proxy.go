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
	"github.com/asaskevich/govalidator"
	"github.com/chatgpt-accesstoken/render"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s Server) handlerGetProxy(ctx *gin.Context) {
	list, err := s.proxySvc.List(ctx)
	if err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}

	render.JSON(ctx.Writer, list, http.StatusOK)
}

type proxyRequest struct {
	Proxy string `json:"proxy"`
}

func (s Server) handlerPostProxy(ctx *gin.Context) {
	in := new(proxyRequest)
	if err := ctx.BindJSON(in); err != nil {
		render.BadRequest(ctx.Writer, err)
		return
	}

	if govalidator.IsNull(in.Proxy) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot find proxy"))
		return
	}

	if err := s.proxySvc.Add(ctx, in.Proxy); err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (s Server) handlerDeleteProxy(ctx *gin.Context) {
	ip := ctx.Param("ip")
	if govalidator.IsNull(ip) {
		render.BadRequest(ctx.Writer, errors.New("api: cannot find ip"))
		return
	}

	if err := s.proxySvc.Delete(ctx, ip); err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
}
