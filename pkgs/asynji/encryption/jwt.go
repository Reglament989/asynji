package encryption

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userid string, refresh bool) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["userId"] = userid
	if refresh {
		atClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	} else {
		atClaims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte("SUBKA"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("SUBKA"), nil
	})

	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return claims["userId"].(string), nil
	}
	return "", err
}

func CreateTokenPair(payload string) (string, string, error) {
	token, err := CreateToken(payload, false)
	if err != nil {
		return "", "", err
	}
	refresh, _ := CreateToken(payload, true)
	return token, refresh, nil
}
