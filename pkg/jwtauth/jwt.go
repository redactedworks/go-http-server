package jwtauth

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	jwt.Claims
}
