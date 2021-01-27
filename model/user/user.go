package user

type (
	// 账号信息
	Account struct {
		ID          int    `json:"id"`       // 账号 id
		Token       string `json:"token"`    // 令牌(唯一)
		Account     string `json:"account"`  // 登录账号名
		Password    string `json:"password"` // 登录密码
		AccountType int    `json:"type"`     // 账号类型 0-游客 1-绑定用户
		Platform    int    `json:"platform"` // 所属平台
	}

	// 用户
	User struct {
		Account Account `json:"account"` // 账号信息
	}
)

// 新的用户，没有任何信息
func FreshUser() User {
	return User{
		Account: Account{},
	}
}
