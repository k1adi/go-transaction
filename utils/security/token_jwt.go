package security

import (
	"fmt"
	"go-transaction/config"
	"go-transaction/model"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(auth model.Auth, role string) string {
	cfg := config.NewConfig()

	now := time.Now().UTC()
	end := now.Add(cfg.AccessTokenLifeTime)

	claims := &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Id:       auth.Id,
		Username: auth.Username,
		Role:     role,
	}

	token := jwt.NewWithClaims(cfg.JwtSigningMethod, claims)
	ss, err := token.SignedString(cfg.JwtSignatureKey)
	if err != nil {
		log.Printf("failed to create access token : %s", err.Error())

	}
	return ss
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	cfg := config.NewConfig()
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != cfg.JwtSigningMethod {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid parse token : %s", err.Error())
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("invalid token mapclaims")
	}
	return claims, nil
}
