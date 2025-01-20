package services

import (
	"errors"
	"github.com/google/uuid"
	"goim/dao"
	"goim/model"
	"goim/utils/crypto"
	"goim/utils/jwt"
	"goim/utils/redi"
	"strings"
	"time"
)

const (
	UserExisted   = "用户已存在"
	UserLoginErr  = "账号或密码错误"
	RefreshFailed = "该令牌在10分钟内已被刷新，无法再次刷新"
)

// Register 注册
func Register(user model.User) error {
	existed, err := dao.IsExistedUser(user.UserName)
	if err != nil {
		return errors.New("查询用户失败")
	}
	if existed {
		return errors.New("用户已存在")
	}

	// 密码加密
	HashPassword := crypto.CryptoPwd(user.PassWord)

	user.Uuid = uuid.New().String()
	user.PassWord = HashPassword

	err = dao.CreateUser(&user)
	if err != nil {
		return errors.New("创建用户失败")
	}
	return nil
}

// Login 登录
func Login(username, password string) (token *jwt.Token, err error) {
	User, Id, err := Verify(username, password)
	if err != nil {
		return nil, errors.New("账号或密码错误")
	}

	token, err = jwt.CreateToken(Id, User.Uuid, User.UserName)
	if err != nil {
		return nil, errors.New("生成 token 失败")
	}
	return token, nil
}

// Verify 核实账号密码
func Verify(username, password string) (user model.User, Id uint, err error) {
	U, err := dao.GetUserByUsername(username)
	if err != nil {
		return model.User{}, 0, errors.New("用户不存在")
	}

	// 密码加密
	HashPassword := crypto.CryptoPwd(password)
	User, err := dao.CheckUser(username, HashPassword)
	if err != nil {
		return model.User{}, 0, errors.New("账号或密码错误")
	}

	user = model.User{
		UserName: User.UserName,
		PassWord: User.PassWord,
		Uuid:     User.Uuid,
		Avatar:   User.Avatar,
		Email:    User.Email,
		NickName: User.NickName,
	}

	return user, U.ID, nil
}

// RefreshToken 刷新token
func RefreshToken(refreshToken string) (newToken *jwt.Token, err error) {
	parts := strings.Split(refreshToken, ".")
	payload := parts[1] // 获取第一个 . 后面的部分 (Payload)

	blockKey := payload[:len(payload)/2] + "-BLOCK"
	blocked := redi.Get(blockKey)
	if blocked != "" {
		lockTime, err := time.Parse(time.RFC3339, blocked)
		if err == nil && time.Since(lockTime) < 10*time.Minute {
			return nil, errors.New("该令牌在10分钟内已被刷新，无法再次刷新")
		}
	}

	newToken, err = jwt.RefreshAccessToken(refreshToken)

	// 设置锁时间为当前时间
	redi.Set(blockKey, time.Now().Format(time.RFC3339))

	return newToken, nil
}
