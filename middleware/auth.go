package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goim/utils/jwt"
	"net/http"
	"strings"
)

// JwtMiddleware JWT中间件，用于解析和验证JWT
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求中的Authorization头部
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort() // 请求被中断
			return
		}

		// Authorization头部格式为: "Bearer <token>"
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization format"})
			c.Abort() // 请求被中断
			return
		}

		// 提取Token
		token := authParts[1]

		// 验证Token，假设我们正在验证access token
		payload, err := jwt.ValidateToken(token, false) // 如果是刷新令牌，传true
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Token validation failed: %v", err)})
			c.Abort() // 请求被中断
			return
		}

		// 将解析出来的用户信息附加到上下文中，供下游使用
		c.Set("username", payload.UserName) // Gin的上下文传递数据，后续处理中可以通过c.Get("user")获取

		// 继续处理
		c.Next()
	}
}
