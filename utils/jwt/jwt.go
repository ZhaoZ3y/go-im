package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Payload 载荷
type Payload struct {
	UUid      string `json:"UUid"`
	UserName  string `json:"UserName"`
	Exp       int64  `json:"Exp"`
	Iss       int64  `json:"Iss"`
	TokenType string `json:"TokenType"`
}

// Token 令牌结构体
type Token struct {
	RefreshToken string `json:"RefreshToken"`
	AccessToken  string `json:"AccessToken"`
}

// 定义常量
const (
	TokenDuration    = 3 * 24 * time.Hour
	RefreshDuration  = 15 * 24 * time.Hour
	SecretKey        = "ST20EW1YSB14MHW103YOUJFN153IM3W59Q"
	RefreshSecretKey = "RFSTEL25YY18SW10WBY3YING46GOIM686Z"
)

// 黑名单存储
var blackList = struct {
	sync.RWMutex
	tokens map[string]struct{}
}{
	tokens: make(map[string]struct{}),
}

// CreateToken 生成token结构体
// CreateToken 生成token结构体
func CreateToken(uuid, username string) (*Token, error) {
	// 生成 access token
	accessToken, err := generateToken(uuid, username, TokenDuration, SecretKey, false)
	if err != nil {
		return nil, fmt.Errorf("生成 access token 失败: %v", err)
	}
	if accessToken == "" {
		return nil, errors.New("生成的 access token 为空")
	}

	// 生成简短的 refresh token
	refreshToken, err := generateToken(uuid, username, RefreshDuration, RefreshSecretKey, true)
	if err != nil {
		return nil, fmt.Errorf("生成 refresh token 失败: %v", err)
	}
	if refreshToken == "" {
		return nil, errors.New("生成的 refresh token 为空")
	}

	token := &Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// 确保 token 不是空值
	if token.AccessToken == "" || token.RefreshToken == "" {
		return nil, errors.New("生成的 token 为空")
	}

	return token, nil
}

// generateToken 生成JWT token
func generateToken(uuid, username string, duration time.Duration, secretKey string, isRefresh bool) (string, error) {
	header := map[string]string{
		"Alg": "HS256",
		"Typ": "JWT",
	}

	payload := Payload{
		UUid:      uuid + strconv.Itoa(rand.Intn(math.MaxInt)),
		UserName:  username,
		Exp:       time.Now().Add(duration).Unix(),
		Iss:       time.Now().Unix(),
		TokenType: "Bearer Token",
	}

	// 针对RefreshToken简化payload
	if isRefresh {
		// 仅保留必要字段，减少token长度
		payload = Payload{
			UUid: uuid + strconv.Itoa(rand.Intn(math.MaxInt)), // 保留用户ID
			Exp:  payload.Exp,                                 // 保留过期时间
		}
	}

	headerJson, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerEncode := base64.RawURLEncoding.EncodeToString(headerJson)

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	payloadEncode := base64.RawURLEncoding.EncodeToString(payloadJson)

	signature := generateSignature(headerEncode+"."+payloadEncode, secretKey)
	return headerEncode + "." + payloadEncode + "." + signature, nil
}

// generateSignature 生成签名
func generateSignature(content, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(content))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

// ValidateToken 验证令牌
func ValidateToken(token string, isRefresh bool) (*Payload, error) {
	key := SecretKey
	if isRefresh {
		key = RefreshSecretKey
	}

	// 检查是否在黑名单中
	blackList.RLock()
	if _, found := blackList.tokens[token]; found {
		blackList.RUnlock()
		return nil, errors.New("该令牌已失效")
	}
	blackList.RUnlock()

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("无效的令牌格式")
	}

	expectedSignature := generateSignature(parts[0]+"."+parts[1], key)
	if expectedSignature != parts[2] {
		return nil, errors.New("无效的签名")
	}

	payloadJson, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("解码payload失败: %v", err)
	}

	var claims Payload
	if err := json.Unmarshal(payloadJson, &claims); err != nil {
		return nil, fmt.Errorf("解析claims失败: %v", err)
	}

	if time.Now().Unix() > claims.Exp {
		return nil, errors.New("令牌已过期")
	}

	return &claims, nil
}

// RefreshAccessToken 使用刷新令牌生成新的访问令牌
func RefreshAccessToken(refreshToken string) (*Token, error) {
	// 验证刷新令牌
	claims, err := ValidateToken(refreshToken, true)
	if err != nil {
		return nil, fmt.Errorf("刷新令牌验证失败: %v", err)
	}
	return CreateToken(claims.UUid, claims.UserName)
}
