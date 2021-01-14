package services

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv"
)

//var jwtKey = []byte(os.Getenv("SECRET_KEY"))
var jwtKey = []byte("qwertyuiopasdfghjklzxcvbnm123456")

// Claims defines jwt claims
type Claims struct {
	UserId string `json:"user_id"`
	CompanyId string `json:"company_id"`
	jwt.StandardClaims
}

// DecodeToken handles decoding a jwt token
func DecodeToken(tkStr string) (UserId string, CompanyId string, err error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "","", err
		}
		return "","", err
	}

	if !tkn.Valid {
		return "","", err
	}

	return claims.UserId, claims.CompanyId, nil
}
