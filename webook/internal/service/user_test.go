package service

import (
	"context"
	"github.com/sgy111222333/basic-go/webook/internal/domain"
	"github.com/sgy111222333/basic-go/webook/internal/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	password := []byte("123456#hello")
	encrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	assert.NoError(t, err)
	println(string(encrypted))
	err = bcrypt.CompareHashAndPassword(encrypted, []byte("123456#hello"))
	assert.NoError(t, err)
}

func Test_userService_Login(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) repository.UserRepository
		// 预期输入
		ctx      context.Context
		email    string
		password string
		// 预期输出
		wantUser domain.User
		wantErr  error
	}{
		{},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := tc.mock(ctrl)
			svc := NewUserService(repo)
			svc.Login(tc.ctx, tc.email, tc.password)
		})
	}
}
