package services

import (
	"serveur/server/models"
)

import (
	"github.com/golang-jwt/jwt"
	_ "github.com/golang-jwt/jwt"
	_ "os"
)

func NewClientAccessToken(claims models.ClientClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte("SECRET"))
}

func NewRestaurantAccessToken(claims models.RestaurantClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte("SECRET"))
}

func ParseAccessToken(accessToken string) *models.ClientClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &models.ClientClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})
	return parsedAccessToken.Claims.(*models.ClientClaims)
}