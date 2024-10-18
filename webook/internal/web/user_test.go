package web

import (
	"bytes"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sgy111222333/basic-go/webook/internal/domain"
	"github.com/sgy111222333/basic-go/webook/internal/service"
	svcmocks "github.com/sgy111222333/basic-go/webook/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_SignUp(t *testing.T) {
	// testCases 里面声明测试用例的格式并填充
	testCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) (service.UserService, service.CodeService)
		// 构造预期中的输入
		reqBuilder func(t *testing.T) *http.Request
		// 预期中的输出
		wantCode int
		wantBody string
	}{ // 这是切片的大括号
		{ // 这是一个用例的大括号
			name: "注册成功",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				userSvc.EXPECT().Signup(gomock.Any(), domain.User{
					Email:    "123@qq.com",
					Password: "Hello#world123",
				}).Return(nil)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				return userSvc, codeSvc
			},
			reqBuilder: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewReader([]byte(`{
					"email": "123@qq.com",
					"password": "Hello#world123",
					"confirmPassword": "Hello#world123"}`)))
				req.Header.Set("Content-Type", "application/json")
				assert.NoError(t, err)
				return req
			},
			wantCode: http.StatusOK,
			wantBody: "注册成功",
		},
		{ // 这是一个用例的大括号
			name: "Bind出错",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				return userSvc, codeSvc
			},
			reqBuilder: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewReader([]byte(`{
					"email": "123@q`)))
				req.Header.Set("Content-Type", "application/json")
				assert.NoError(t, err)
				return req
			},
			wantCode: http.StatusBadRequest,
		},
		{ // 这是一个用例的大括号
			name: "邮箱格式不对",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				return userSvc, codeSvc
			},
			reqBuilder: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewReader([]byte(`{
					"email": "123",
					"password": "Hello#world123",
					"confirmPassword": "Hello#world123"}`)))
				req.Header.Set("Content-Type", "application/json")
				assert.NoError(t, err)
				return req
			},
			wantCode: http.StatusOK,
			wantBody: "邮箱格式不对",
		},
		{ // 这是一个用例的大括号
			name: "两次密码输入不同",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				return userSvc, codeSvc
			},
			reqBuilder: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewReader([]byte(`{
					"email": "123@qq.com",
					"password": "Hello#world123",
					"confirmPassword": "Hello#world1234"}`)))
				req.Header.Set("Content-Type", "application/json")
				assert.NoError(t, err)
				return req
			},
			wantCode: http.StatusOK,
			wantBody: "两次输入密码不一致",
		},
		{ // 这是一个用例的大括号
			name: "密码格式不对",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				return userSvc, codeSvc
			},
			reqBuilder: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewReader([]byte(`{
					"email": "123@qq.com",
					"password": "123",
					"confirmPassword": "123"}`)))
				req.Header.Set("Content-Type", "application/json")
				assert.NoError(t, err)
				return req
			},
			wantCode: http.StatusOK,
			wantBody: "密码格式不对, 必须包含大小写字母和数字的组合，可以使用特殊字符，长度在8-16之间",
		},
		{ // 这是一个用例的大括号
			name: "系统错误",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				userSvc.EXPECT().Signup(gomock.Any(), domain.User{
					Email:    "123@qq.com",
					Password: "Hello#world123",
				}).Return(errors.New("DB出错")) //new任意一个error
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				return userSvc, codeSvc
			},
			reqBuilder: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewReader([]byte(`{
					"email": "123@qq.com",
					"password": "Hello#world123",
					"confirmPassword": "Hello#world123"}`)))
				req.Header.Set("Content-Type", "application/json")
				assert.NoError(t, err)
				return req
			},
			wantCode: http.StatusOK,
			wantBody: "系统错误",
		},
		{ // 这是一个用例的大括号
			name: "邮箱冲突",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				userSvc.EXPECT().Signup(gomock.Any(), domain.User{
					Email:    "123@qq.com",
					Password: "Hello#world123",
				}).Return(service.ErrDuplicateEmail)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				return userSvc, codeSvc
			},
			reqBuilder: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewReader([]byte(`{
					"email": "123@qq.com",
					"password": "Hello#world123",
					"confirmPassword": "Hello#world123"}`)))
				req.Header.Set("Content-Type", "application/json")
				assert.NoError(t, err)
				return req
			},
			wantCode: http.StatusOK,
			wantBody: "该邮箱已被注册",
		},
	}

	for _, tc := range testCases {
		// 前面的t和后面的t不是同一个
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish() //新版本go mock可以不写这句
			// 构造 handler
			userSvc, codeSvc := tc.mock(ctrl)
			hdl := NewUserHandler(userSvc, codeSvc)
			// 准备服务器, 注册路由
			server := gin.Default()
			hdl.RegisterRoutes(server)
			// 准备Req和记录的 recorder
			req := tc.reqBuilder(t)
			recorder := httptest.NewRecorder()
			// 执行
			server.ServeHTTP(recorder, req)
			// 断言结果
			assert.Equal(t, tc.wantCode, recorder.Code)
			assert.Equal(t, tc.wantBody, recorder.Body.String())
		})
	}
}

func TestUserEmailPattern(t *testing.T) {
	// Table Driven 模式, 这是一个匿名结构体
	testCases := []struct {
		name  string
		email string
		match bool
	}{
		{
			name:  "不带@",
			email: "123456",
			match: false,
		},
		{
			name:  "带@, 没后缀",
			email: "123456@",
			match: false,
		}, {
			name:  "合法邮箱",
			email: "123456@qq.com",
			match: true,
		},
	}
	h := NewUserHandler(nil, nil)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			match, err := h.emailRexExp.MatchString(tc.email)
			require.NoError(t, err)
			assert.Equal(t, tc.match, match)
		})
	}
}

//func TestHTTP(t *testing.T) {
//	req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewReader([]byte("我的请求体")))
//	assert.NoError(t, err)
//	recorder := httptest.NewRecorder()
//	assert.Equal(t, http.StatusOK, recorder.Code)
//}

func TestMock(t *testing.T) {
	// 必须先声明EXPECT再调用mock的方法
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() //新版本go mock可以不写这句
	userSvc := svcmocks.NewMockUserService(ctrl)
	// 调用userSvc后的预期: 第一个参数无所谓, 第二个必须和这里设置的模拟场景相同
	// 意思是如果signup方法收到了any和内容如下的user, 就返回nil
	userSvc.EXPECT().Signup(gomock.Any(), domain.User{
		Id:    1,
		Email: "123@qq.com",
	}).Return(errors.New("DB出错")) // return nil是正常, return error是模拟出错

	err := userSvc.Signup(context.Background(), domain.User{
		Id:    1,
		Email: "123@qq.com",
	})
	t.Log(err)
}
