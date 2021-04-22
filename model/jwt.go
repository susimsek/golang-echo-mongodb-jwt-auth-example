package model

import "github.com/dgrijalva/jwt-go"

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	ID   string `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	jwt.StandardClaims
}

type Token struct {
	Token string `json:"token" xml:"token"`
}
