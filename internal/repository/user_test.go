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
	"github.com/gelleson/assistant/internal/auth"
	"github.com/gelleson/assistant/internal/model"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
	"testing"
)

func genLongPassword() string {
	var password strings.Builder

	for i := 0; i < 1000; i++ {
		password.WriteString("TTETEEEEEEADA")
	}

	return password.String()
}

type UserRepositoryTestSuite struct {
	suite.Suite
	repo *UserRepository
}

func (r *UserRepositoryTestSuite) SetupTest() {
	dialect := sqlite.Open(":memory:")
	db, _ := gorm.Open(dialect, &gorm.Config{})

	_ = db.AutoMigrate(&model.User{})

	r.repo = &UserRepository{
		db: db,
	}

	password, _ := auth.HashPassword("password")

	_ = db.Create(&model.User{
		Username: "test_1",
		Email:    "test@mail.tester",
		Password: password,
	})
}

func (r UserRepositoryTestSuite) TestGetByUsername() {
	user, err := r.repo.GetUserByUsername("test_1")

	r.Assert().Nil(err)
	r.NotEmpty(user)

	user, err = r.repo.GetUserByUsername("test_2")

	r.Assert().Error(err)
	r.Empty(user)

}

func (r UserRepositoryTestSuite) TestGetByID() {
	user, err := r.repo.GetUserById(1)

	r.Assert().Nil(err)
	r.NotEmpty(user)

	user, err = r.repo.GetUserById(2)

	r.Assert().Error(err)
	r.Empty(user)
}

func (r UserRepositoryTestSuite) TestDeleteUser() {
	err := r.repo.Delete(1)

	r.Assert().Nil(err)

	err = r.repo.Delete(2)

	r.Assert().Nil(err)
}

func (r UserRepositoryTestSuite) TestUpdate() {

	user := model.User{
		FirstName: "TEST_NAME",
		LastName:  "TEST_LASTNAME",
	}
	updatedUser, err := r.repo.Update(2, user)

	r.Assert().Error(err)

	updatedUser, err = r.repo.Update(1, user)

	r.Assert().Nil(err)
	r.Assert().NotEmpty(updatedUser)
	r.Assert().Equal("TEST_NAME", updatedUser.FirstName)
	r.Assert().Equal("TEST_LASTNAME", updatedUser.LastName)
}

func (r UserRepositoryTestSuite) TestExistUsername() {
	isExist := r.repo.ExistUsername("test_1")

	r.Assert().True(isExist)

	isExist = r.repo.ExistUsername("test_2")

	r.Assert().False(isExist)
}

func (r UserRepositoryTestSuite) TestExistEmail() {
	isExist := r.repo.ExistEmail("test@mail.tester")

	r.Assert().True(isExist)

	isExist = r.repo.ExistEmail("test2@mail.tester")

	r.Assert().False(isExist)
}

func (r UserRepositoryTestSuite) TestLastActive() {

	r.repo.LastActiveUpdate(1)

	var user model.User

	r.repo.db.First(&user, 1)

	r.Assert().NotEmpty(user.LastActive)
}

func (r UserRepositoryTestSuite) TestCreate() {
	empty, err := r.repo.Create(model.User{})

	r.Assert().Error(err)
	r.Assert().Empty(empty)

	empty, err = r.repo.Create(model.User{
		Username: "test_1",
		Email:    "test@mail.tester",
		Password: "password",
	})

	r.Assert().Error(err)
	r.Assert().Empty(empty)

	empty, err = r.repo.Create(model.User{
		Username: "test_2",
		Email:    "test@mail.tester",
		Password: "password",
	})

	r.Assert().Error(err)
	r.Assert().Empty(empty)

	user, err := r.repo.Create(model.User{
		Username: "test_2",
		Email:    "test2@mail.tester",
		Password: "password",
	})

	r.Assert().Nil(err)
	r.Assert().NotEmpty(user)
	r.Assert().Equal("test_2", user.Username)
}

func TestRunUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
