package utils

import "github.com/golang-jwt/jwt/v5"

func GetJwtToken(secretKey string, iat, seconds, payload int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func SetJwtToken(secretKey string, token string) error {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return err
	}
	return nil
}
