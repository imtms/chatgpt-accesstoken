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

func (s Server) handlerGetFakeopen(ctx *gin.Context) {
	URL, err := s.fakeopenStore.Get()
	if err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}
	ctx.JSON(200, gin.H{
		"URL": URL,
	})
}

type URLRequest struct {
	URL string `json:"URL"`
}

func (s Server) hanlderPostFakeopen(ctx *gin.Context) {
	in := new(URLRequest)
	if err := ctx.BindJSON(in); err != nil {
		render.BadRequest(ctx.Writer, err)
		return
	}
	if govalidator.IsNull(in.URL) {
		render.BadRequest(ctx.Writer, errors.New("api: URL can't be empty"))
		return
	}
	if err := s.fakeopenStore.Set(in.URL); err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (s Server) handlerDeleteFakeopen(ctx *gin.Context) {
	if err := s.fakeopenStore.Delete(); err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (s Server) hanlderPutFakeopen(ctx *gin.Context) {
	in := new(URLRequest)
	if err := ctx.BindJSON(in); err != nil {
		render.BadRequest(ctx.Writer, err)
		return
	}
	if govalidator.IsNull(in.URL) {
		render.BadRequest(ctx.Writer, errors.New("api: URL can't be empty"))
		return
	}
	if err := s.fakeopenStore.Set(in.URL); err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}
