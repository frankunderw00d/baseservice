package user

import (
	"context"
	"jarvis/base/database"
)

type (
	// 账号信息 `dynamic_account`
	Account struct {
		ID          int    `json:"id"`       // 账号 id
		Token       string `json:"token"`    // 令牌(唯一)
		Account     string `json:"account"`  // 登录账号名
		Password    string `json:"password"` // 登录密码
		AccountType int    `json:"type"`     // 账号类型 0-游客 1-绑定用户
		Platform    int    `json:"platform"` // 所属平台
	}

	// 用户信息 `dynamic_userInfo`
	Info struct {
		ID                int    `json:"id"`                   // 用户 id
		AccountToken      string `json:"account_token"`        // 账号唯一标识
		Name              string `json:"name"`                 // 用户名
		Age               int    `json:"age"`                  // 用户年龄
		Sex               bool   `json:"sex"`                  // 用户性别
		HeadImage         int    `json:"head_image"`           // 用户头像序号
		Vip               int    `json:"vip"`                  // 用户 vip 等级
		GameBgMusicVolume int    `json:"game_bg_music_volume"` // 背景音乐音量
		GameEffectVolume  int    `json:"game_effect_volume"`   // 音效音量
		AccountBalance    int64  `json:"account_balance"`      // 账户余额(单位:分)
	}

	// 用户
	User struct {
		Account Account `json:"account"` // 账号信息
		Info    Info    `json:"info"`    // 用户信息
	}
)

// 新的用户，没有任何信息
func FreshUser() User {
	return User{
		Account: Account{},
		Info:    Info{},
	}
}

// 通过账号密码加载用户信息
func (u *User) LoadInfoByAccountAndPassword(account, password string) error {
	if err := u.Account.loadByAccountAndPassword(account, password); err != nil {
		return err
	}

	if err := u.Info.loadByAccountToken(u.Account.Token); err != nil {
		return err
	}

	return nil
}

// 通过账号密码加载用户信息
func (u *User) LoadInfoByToken(token string) error {
	if err := u.Account.loadByToken(token); err != nil {
		return err
	}

	if err := u.Info.loadByAccountToken(token); err != nil {
		return err
	}

	return nil
}

// 存储信息
func (u *User) Store() error {
	if err := u.Account.Store(); err != nil {
		return err
	}

	if err := u.Info.Store(u.Account.Token); err != nil {
		return err
	}

	return nil
}

// 通过账号密码加载账号信息
func (a *Account) loadByAccountAndPassword(account, password string) error {
	// 获取 MySQL 连接
	mysqlConn, err := database.GetMySQLConn()
	if err != nil {
		return err
	}
	defer mysqlConn.Close()

	row := mysqlConn.QueryRowContext(context.Background(), "select * from `jarvis`.`dynamic_account` where account = ? and password = ?", account, password)
	if err := row.Scan(&a.ID, &a.Token, &a.Account, &a.Password, &a.AccountType, &a.Platform); err != nil {
		return err
	}

	return nil
}

// 通过 token 加载账号信息
func (a *Account) loadByToken(token string) error {
	// 获取 MySQL 连接
	mysqlConn, err := database.GetMySQLConn()
	if err != nil {
		return err
	}
	defer mysqlConn.Close()

	row := mysqlConn.QueryRowContext(context.Background(), "select * from `jarvis`.`dynamic_account` where token = ?", token)
	if err := row.Scan(&a.ID, &a.Token, &a.Account, &a.Password, &a.AccountType, &a.Platform); err != nil {
		return err
	}

	return nil
}

// 存储信息
func (a *Account) Store() error {
	// 获取 MySQL 连接
	mysqlConn, err := database.GetMySQLConn()
	if err != nil {
		return err
	}
	defer mysqlConn.Close()

	_, err = mysqlConn.ExecContext(context.Background(), "insert into `jarvis`.`dynamic_account`(token, account, password, type, platform) values (?,?,?,?,?)",
		a.Token,
		a.Account,
		a.Password,
		a.AccountType,
		a.Platform)

	return err
}

// 通过 token 加载账号信息
func (i *Info) loadByAccountToken(token string) error {
	// 获取 MySQL 连接
	mysqlConn, err := database.GetMySQLConn()
	if err != nil {
		return err
	}
	defer mysqlConn.Close()

	row := mysqlConn.QueryRowContext(context.Background(), "select id, account_token, name, age, sex, head_image, vip, game_bg_music_volume, game_effect_volume, account_balance from `jarvis`.`dynamic_userInfo` where account_token = ?", token)
	if err := row.Scan(
		&i.ID,
		&i.AccountToken,
		&i.Name,
		&i.Age,
		&i.Sex,
		&i.HeadImage,
		&i.Vip,
		&i.GameBgMusicVolume,
		&i.GameEffectVolume,
		&i.AccountBalance,
	); err != nil {
		return err
	}

	return nil
}

// 通过 token 加载账号信息
func (i *Info) Store(token string) error {
	// 获取 MySQL 连接
	mysqlConn, err := database.GetMySQLConn()
	if err != nil {
		return err
	}
	defer mysqlConn.Close()

	_, err = mysqlConn.ExecContext(context.Background(), "insert into `jarvis`.`dynamic_userInfo`(account_token, name) values (?,?)",
		token,
		i.Name)

	return err
}

// 通过 token 修改账号信息
func (i *Info) Update() error {
	// 获取 MySQL 连接
	mysqlConn, err := database.GetMySQLConn()
	if err != nil {
		return err
	}
	defer mysqlConn.Close()

	_, err = mysqlConn.ExecContext(context.Background(), "update `jarvis`.`dynamic_userInfo` set name = ? , age = ? , sex = ? , head_image = ? , game_bg_music_volume = ? , game_effect_volume = ? where account_token = ?",
		i.Name,
		i.Age,
		i.Sex,
		i.HeadImage,
		i.GameBgMusicVolume,
		i.GameEffectVolume,
		i.AccountToken)

	return err
}

// 通过 token 修改账号信息
func (i *Info) UpdateAccountBalance() error {
	// 获取 MySQL 连接
	mysqlConn, err := database.GetMySQLConn()
	if err != nil {
		return err
	}
	defer mysqlConn.Close()

	_, err = mysqlConn.ExecContext(context.Background(), "update `jarvis`.`dynamic_userInfo` set account_balance = ? where account_token = ?",
		i.AccountBalance,
		i.AccountToken)

	return err
}
