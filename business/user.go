package business

import "strconv"

func ValidateToken(userId, shopId, token string) bool {
	// todo: 校验登录态
	if u, _ := strconv.Atoi(userId); u > 10 {
		return false
	}
	return true
}
