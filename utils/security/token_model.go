package security

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	jwt.RegisteredClaims
	Id       string
	Username string
	Role     string
}
