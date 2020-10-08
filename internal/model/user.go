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

package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	// gorm.Model pk
	gorm.Model
	// Username is required
	Username string `json:"username" validate:"required" gorm:"unique,index"`
	// Email to contact with user
	Email string `json:"email" validate:"required,email" gorm:"unique,index"`
	// FirstName of the user
	FirstName string `json:"firstName"`
	// LastName of the user
	LastName string `json:"lastName"`
	// IsAdmin part of admin team
	IsAdmin bool `json:"isAdmin"`
	// Password of the user encrypted
	Password string `json:"-" validate:"required"`
	// LastActive is time when used api
	LastActive time.Time `json:"lastActive"`
}