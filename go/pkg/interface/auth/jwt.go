package auth

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
)

func CreateToken(userID string) (string, error) {
	// tokenの作成
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// claimsの設定
	token.Claims = jwt.MapClaims{
		"user": userID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	// 署名
	var secretKey = "l-semi-chat"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte("secret")
		return b, nil
	})
	return token, err
}
