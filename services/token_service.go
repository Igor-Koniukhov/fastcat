package services

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(userID, lifetimeMinutes int, secret string) (string, error) {
	claims := &JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifetimeMinutes)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString, secret string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}
	return claims, nil
}

func GetTokenFromBearerString(input string) (string, error) {
	if input == "" {
		return "", errors.New("no authorization header received")
	}
	headerParts := strings.Split(input, "Bearer=")
	if len(headerParts) != 2  {
		return "", errors.New("no authorization header received")
	}
	token := strings.TrimSpace(headerParts[1])
	if len(token) == 0 {
		return "", errors.New("authorization token wrong size")
	}
	return token, nil
}

func IsAuthenticated(ctx context.Context) (string, bool)  {
	exists, ok := ctx.Value("user_id").(string)
	return exists, ok
}