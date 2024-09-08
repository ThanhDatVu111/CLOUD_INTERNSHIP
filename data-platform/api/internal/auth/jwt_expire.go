package auth

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetJwtDuration(jwtExpiryStr string, jwtExpiryDuration time.Duration) (*time.Duration, error) {
	jwtExpiry, err := strconv.Atoi(jwtExpiryStr)
	if err != nil {
		log.Println("Invalid JWT expiry duration in .env file")
		return nil, err
	}
	jwtExpiryDuration = time.Duration(jwtExpiry) * time.Minute
	return &jwtExpiryDuration, nil
}

var secretKey = "your_secret_key"

func createTestToken(expiryDuration time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiryDuration).Unix()

	claims := &jwt.MapClaims{
		"exp": expirationTime,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJwtExpiry(tokenString string, tokenIssueTime time.Time) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0)
			return expirationTime.After(tokenIssueTime), nil
		}
		return false, fmt.Errorf("exp claim not found in token")
	}
	return false, fmt.Errorf("invalid token")
}
