package authorization

// Invalid 检查用户认证信息是否有效
//
// Returns:
//
// - bool: true: 无效 false: 有效
func (auth *Authorization) Invalid() bool {
	if auth.Missing {
		return true
	}
	if auth.UserId < 1 {
		return true
	}

	return false
}
