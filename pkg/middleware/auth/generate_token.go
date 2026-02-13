package auth

import (
	"time"
	"url-shortening-service/pkg/domain"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userEmail, password string) (domain.TokenResponse, error) {

	ctg := LoadJWTCtg()

	secret := ctg.Secret

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": userEmail,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secret)

	Claims := token.Claims.(jwt.MapClaims)
	uName := Claims["email"].(string)

	if err != nil {
		return domain.TokenResponse{}, err
	}

	return domain.TokenResponse{
		Token: tokenString,
		Email: uName,
	}, nil

}
