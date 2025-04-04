package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_GetIdByToken(t *testing.T) {
	t.Run("当token存在时应该返回对应的ID", shouldReturnIdWhenTokenExist)
	t.Run("当token不存在时应该返回0", shouldReturnZeroWhenTokenNotExist)
}

func shouldReturnIdWhenTokenExist(t *testing.T) {
	testCases := []struct {
		token    string
		expected int64
		message  string
	}{
		{"a", 1, "token 'a' 应该返回用户ID 1"},
		{"b", 2, "token 'b' 应该返回用户ID 2"},
	}

	repo := Inst()
	for _, tc := range testCases {
		id := repo.GetIdByToken(tc.token)
		assert.Equal(t, tc.expected, id, tc.message)
	}
}

func shouldReturnZeroWhenTokenNotExist(t *testing.T) {
	id := Inst().GetIdByToken("nonexistent_token")
	assert.Equal(t, int64(0), id, "不存在的token应该返回0")
}
