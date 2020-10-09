/*
 * MIT License
 *
 * Copyright (c) 2020 gelleson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package middleware

import (
	"github.com/gelleson/assistant/internal/model"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"net/http"
	"strings"
)

func IsAuthorized(skipUrls ...string) func(ctx *macaron.Context, sess session.Store) {

	return func(ctx *macaron.Context, sess session.Store) {

		for _, url := range skipUrls {

			if strings.Contains(ctx.Req.RequestURI, url) {
				ctx.Next()
				return
			}

		}

		if user := sess.Get("user"); user == nil {
			ctx.Error(http.StatusForbidden, "user not authorized")
			return
		}

		ctx.Data["user"] = sess.Get("user")

		ctx.Next()
	}
}

func IsAdmin(ctx *macaron.Context, sess session.Store) {

	if user, exist := ctx.Data["user"]; !exist && !user.(model.User).IsAdmin {
		ctx.Error(http.StatusForbidden, "user have not enough permissions")
		return
	}

	ctx.Next()
}
