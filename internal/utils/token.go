package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtSecret = []byte("fsdgf2343543ttgfhgffgh54343343gthsdfSEWFG")

func GenerateJWT(userID int, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // add 24hour
		"iss":     "rekeningku",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func DecryptJWT(tokenString string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate signed token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// validate token
	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	// convert claims to mapClaims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to parse claims")
	}

	if exp, ok := claims["exp"].(float64); ok {
		expirationTime := time.Unix(int64(exp), 0)
		if time.Now().After(expirationTime) {
			return nil, errors.New("token has expired")
		}
	}

	return claims, nil
}
