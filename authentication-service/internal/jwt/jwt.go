package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(userID int) (common.Token, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	standardClaims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	claims := &common.Claims{
		UserID:         userID,
		StandardClaims: standardClaims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(os.Getenv("SECRET_KEY"))
	if err != nil {
		return common.Token{}, err
	}

	return common.Token{
		Token: tokenString,
	}, nil
}

func ValidateToken(token common.Token) error {
	parsedToken, err := jwt.Parse(token.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return errors.New("token is expired")
		}
	} else {
		return errors.New("invalid token claims")
	}

	return nil
}
