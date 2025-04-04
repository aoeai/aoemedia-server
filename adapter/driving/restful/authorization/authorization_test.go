package authorization

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewAuth(t *testing.T) {
	tests := []struct {
		name            string
		authToken       string
		expectedToken   string
		expectedMissing bool
		expectedUserId  int64
	}{
		{
			name:            "Header中不含认证信息",
			authToken:       "",
			expectedToken:   "",
			expectedMissing: true,
			expectedUserId:  0,
		},
		{
			name:            "Header中包含无效token",
			authToken:       "invalid_token",
			expectedToken:   "invalid_token",
			expectedMissing: false,
			expectedUserId:  0,
		},
		{
			name:            "Header中包含有效token",
			authToken:       "a",
			expectedToken:   "a",
			expectedMissing: false,
			expectedUserId:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建一个新的gin context
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

			// 设置Authorization header
			if tt.authToken != "" {
				c.Request.Header.Set("Authorization", tt.authToken)
			}

			// 调用被测试的函数
			result := NewAuth(c)

			// 验证结果
			assert.Equal(t, tt.expectedToken, result.Token)
			assert.Equal(t, tt.expectedMissing, result.Missing)
			assert.Equal(t, tt.expectedUserId, result.UserId)
		})
	}
}
