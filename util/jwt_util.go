package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	jwtSecret          = "vuniversity"      // JWT密钥
	TokenIssuer        = "vuniversity"      // 签发者
	AccessTokenExpire  = 1 * time.Hour      // Access Token 过期时间
	RefreshTokenExpire = 7 * 24 * time.Hour // Refresh Token 过期时间
)

// Claims 存储 JWT 中的用户信息
type Claims struct {
	UserID int `json:"user_id"` // 用户ID
	jwt.RegisteredClaims
}

// GenerateToken 生成 Token
func GenerateToken(ctx context.Context, userID int, expire time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
			Issuer:    TokenIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// ParseToken 解析 Token
func ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
