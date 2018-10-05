package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim model
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
