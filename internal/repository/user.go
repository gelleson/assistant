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

package repository

import (
	"fmt"
	"github.com/gelleson/assistant/internal/auth"
	"github.com/gelleson/assistant/internal/model"
	"github.com/go-playground/validator"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

// Repository impl for the user
type UserRepository struct {
	db *gorm.DB
}

func (ur UserRepository) GetUserByUsername(username string) (model.User, error) {
	user := model.User{}

	if trn := ur.db.Where("username = ?", username).First(&user); trn.Error != nil {
		return model.User{}, trn.Error
	}

	return user, nil
}

func (ur UserRepository) GetUserById(id int) (model.User, error) {
	user := model.User{}

	if trn := ur.db.First(&user, id); trn.Error != nil {
		return model.User{}, trn.Error
	}

	return user, nil
}

func (ur UserRepository) Create(user model.User) (model.User, error) {

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		return model.User{}, err
	}

	if exist := ur.ExistUsername(user.Username); exist {
		return model.User{}, errors.New(fmt.Sprintf("%s username already exist", user.Username))
	}

	if exist := ur.ExistEmail(user.Email); exist {
		return model.User{}, errors.New(fmt.Sprintf("%s email already exist", user.Email))
	}

	hashedPassword, err := auth.HashPassword(user.Password)

	if err != nil {
		return model.User{}, err
	}

	user.Password = hashedPassword

	if trn := ur.db.Create(&user); trn.Error != nil {
		return model.User{}, trn.Error
	}

	return user, nil
}

func (ur UserRepository) Update(id int, user model.User) (model.User, error) {
	var updatedUser model.User

	if trx := ur.db.First(&updatedUser, id); trx.Error != nil {
		return model.User{}, trx.Error
	}

	if user.FirstName != "" {
		updatedUser.FirstName = user.FirstName
	}

	if user.LastName != "" {
		updatedUser.LastName = user.LastName
	}

	ur.db.Save(&updatedUser)

	return updatedUser, nil
}

func (ur UserRepository) Delete(id int) error {
	ur.db.Delete(&model.User{}, id)

	return nil
}

func (ur UserRepository) ExistEmail(email string) bool {
	user := model.User{}

	if trx := ur.db.Where("email = ?", email).First(&user); trx.Error != nil {
		return false
	}

	return user.Username != ""
}

func (ur UserRepository) ExistUsername(username string) bool {
	user := model.User{}

	if trx := ur.db.Where("username = ?", username).First(&user); trx.Error != nil {
		return false
	}

	return user.Username != ""
}

func (ur UserRepository) LastActiveUpdate(userId int) {
	ur.db.Model(&model.User{}).
		Where("id = ?", userId).
		Update("last_active", time.Now())
}
