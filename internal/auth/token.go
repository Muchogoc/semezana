package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var (
	SECRET = os.Getenv("SECRET_KEY")
)

type CustomClaims struct {
	jwt.StandardClaims

	UserID string `json:"uid"`
}

func SigningKey() []byte {
	return []byte(SECRET)
}

func CreateToken(uid string) (string, error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  "semezana",
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Id:        uuid.NewString(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "semezana",
			Subject:   "chat",
		},
		UserID: uid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(SigningKey())
}

func VerifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return SigningKey(), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("not a custom claim %T", token.Claims)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return claims, nil
}

func GetUIDFromToken(tokenString string) (string, error) {
	claims, err := VerifyToken(tokenString)
	if err != nil {
		return "", err
	}

	return claims.UserID, nil
}
