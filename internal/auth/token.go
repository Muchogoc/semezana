package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Muchogoc/semezana/dto"
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
			NotBefore: time.Now().Unix(),
			Issuer:    "semezana",
			Subject:   "chat",
		},
		UserID: uid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(SigningKey())
}

func ParseToken(ctx context.Context, tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return SigningKey(), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ValidateToken(ctx context.Context, token *jwt.Token) error {
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return fmt.Errorf("not a custom claim %T", token.Claims)
	}

	if !token.Valid {
		return fmt.Errorf("token is invalid")
	}

	// if !claims.VerifyAudience("", false) {
	// 	return fmt.Errorf("token audience is invalid")
	// }

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token is expired")
	}

	// if !claims.VerifyIssuer("", false) {
	// 	return fmt.Errorf("token issuer is invalid")
	// }

	return nil
}

func GetUIDFromContext(ctx context.Context) (string, error) {
	token, err := TokenFromContext(ctx)
	if err != nil {
		return "", err
	}

	claims := token.Claims.(*CustomClaims)

	return claims.UserID, nil
}

func TokenFromContext(ctx context.Context) (*jwt.Token, error) {
	value := ctx.Value(dto.ContextKeyToken)
	if value == nil {
		return nil, fmt.Errorf("no token in context")
	}

	session, ok := value.(*jwt.Token)
	if !ok {
		return nil, fmt.Errorf("invalid token type in context")
	}

	return session, nil
}

func SetTokenContext(ctx context.Context, token *jwt.Token) context.Context {
	return context.WithValue(ctx, dto.ContextKeyToken, token)
}
