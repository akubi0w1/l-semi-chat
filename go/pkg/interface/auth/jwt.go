package auth

import (
	"l-semi-chat/pkg/domain/logger"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken JWTの作成
func CreateToken(id, userID string) (string, error) {
	// tokenの作成
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// claimsの設定
	token.Claims = jwt.MapClaims{
		"id":      id,
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	// 署名
	var secretKey = "l-semi-chat"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		logger.Error("create token: ", err)
		return "", err
	}
	return tokenString, nil
}

// VerifyToken JWTの検証
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("l-semi-chat"), nil
	})
	if err != nil {
		logger.Warn("verifyToken: ", err)
		return token, err
	}
	return token, nil
}
