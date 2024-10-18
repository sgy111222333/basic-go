package web

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserEmailPAttern(t *testing.T) {
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

func TestHTTP(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/users/signup", bytes.NewReader([]byte("MyBody")))
	assert.NoError(t, err)
	recorder := httptest.NewRecorder()
	assert.Equal(t, http.StatusOK, recorder.Code)
}
