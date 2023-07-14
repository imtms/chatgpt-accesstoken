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

func (s Server) handlerGetToken(ctx *gin.Context) {
	list, err := s.akStore.List(ctx)
	if err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}
	render.JSON(ctx.Writer, list, http.StatusOK)
}

func (s Server) handlerDeleteToken(ctx *gin.Context) {
	email := ctx.Query("email")
	if govalidator.IsNull(email) {
		render.BadRequest(ctx.Writer, errors.New("api: email is empty"))
		return
	}

	if err := s.akStore.Delete(ctx, email); err != nil {
		render.InternalError(ctx.Writer, err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}
