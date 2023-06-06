package security

import "github.com/golang-jwt/jwt/v5"

type Jwt struct {
	Claims        jwt.Claims
	SigningMethod jwt.SigningMethod
	SigningKey    string
}

func (j *Jwt) SignWithClaims() (string, error) {
	token := jwt.NewWithClaims(j.SigningMethod, j.Claims)
	return token.SignedString(j.SigningKey)
}
