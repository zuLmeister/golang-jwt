package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func getSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func GenerateAccessToken(userID uint) (string, time.Time, error) {
	minStr := os.Getenv("ACCESS_TOKEN_MINUTES")
	minutes, _ := strconv.Atoi(minStr)
	if minutes <= 0 {
		minutes = 15
	}
	exp := time.Now().Add(time.Duration(minutes) * time.Minute)

	claims := JWTCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(getSecret())
	return signed, exp, err
}

func GenerateRefreshToken(userID uint) (string, time.Time, error) {
	daysStr := os.Getenv("REFRESH_TOKEN_DAYS")
	days, _ := strconv.Atoi(daysStr)
	if days <= 0 {
		days = 30
	}
	exp := time.Now().Add(time.Duration(days) * 24 * time.Hour)

	claims := JWTCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(getSecret())
	return signed, exp, err
}

// Parse and validate token; returns claims if valid
func ParseToken(tokenStr string) (*JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return getSecret(), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
