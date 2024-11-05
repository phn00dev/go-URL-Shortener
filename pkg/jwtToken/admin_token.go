package jwttoken

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const AdminSecretKey = "HJGYUDF!DN^Bdd$%asj*_dasdhas$$ash#dasd&%^$@"

// AdminClaims administratiw ulanyjy maglumatlaryny saklaýar
type AdminClaims struct {
	AdminId  int    `json:"admin_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateAdminToken(adminId int, username, email string) (string, error) {

	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &AdminClaims{
		AdminId:  adminId,
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(AdminSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateAdminToken(tokenString string) (*AdminClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(AdminSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AdminClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
