package crypto

import (
	"crypto/sha256"
	"encoding/base64"
)

func CryptoPwd(password string) (CryPassword string) {
	// 计算密码的SHA-256哈希值
	hash := sha256.Sum256([]byte(password))

	// 将哈希值转换为Base64编码字符串
	return base64.StdEncoding.EncodeToString(hash[:])
}
