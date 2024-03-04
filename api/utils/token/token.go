package token

import (
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID int) (string, error) {
	lifetime, err := strconv.Atoi(os.Getenv("TOKEN_LIFETIME"))
	token_lifetime := time.Now().Add(time.Hour * time.Duration(lifetime)).Unix()
	if err != nil {
		return "", err
	}

	claims := &jwt.MapClaims{
		"authorized": true,
		"exp":        token_lifetime,
		"user_id":    strconv.Itoa(userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func TokenValid(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func ExtractToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenID(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID := claims["user_id"].(string)
		return userID, nil
	}
	return "", err
}
