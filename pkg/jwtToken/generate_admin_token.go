package jwttoken

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const SecretAdminKey = "HJGYUDF!DN^Bdd$%asj*_dasdhas$$ash#dasd&%^$@"

type AdminClaims struct {
	ID        int    `json:"admin_id"`
	Username  string `json:"username"`
	AdminRole string `json:"admin_role"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

func GenerateAdminToken(id int, username, email, adminRole string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &AdminClaims{
		ID:        id,
		Username:  username,
		AdminRole: adminRole,
		Email:     email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateAdminToken(tokenString string) (*AdminClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AdminClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
