package jwt

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

var mySingningKey = []byte("secret")

func CreateToken() (string, error) {
	claims := jwt.MapClaims{
		"uid": "111111",
		"exp": time.Now().Add(time.Second * 4).Unix(),
		"iss": "auth.sunsl.net",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	getTokenRemainingValidity(claims["exp"])

	tokenstring, err := token.SignedString(mySingningKey)
	if err != nil {
		return "", err
	}

	return tokenstring, nil
}

func RequireTokenAuthentication(tokenstring string) (bool, int) {
	token, _ := jwt.ParseWithClaims(tokenstring, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return mySingningKey, nil
	})

	flag := getTokenRemainingValidity(token.Claims.(jwt.MapClaims)["exp"])
	return token.Valid, flag
}

func getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			fmt.Println("*")
			return int(remainer.Seconds())
		}
	}
	return -1
}