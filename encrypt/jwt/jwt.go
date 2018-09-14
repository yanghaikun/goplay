package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"log"
	"fmt"
)


func GenJwtTokenForOnlineSupport() (string, error){
	//var jwtSecretKey = "QJZundppmWpyAsqpolsIqZ2KcDC0cZBx2ZVcWcu8dgmMKx1hToHAhoerPRan30fMYxcKLAoMpbW3QC4v7MOYAeQH0o4KzZxmQJDNfFdzBCfFVTkRcwcHXxyoqgkyR3I6ZE0ZTAdT9Q9l1d0We8oJRTFZhFcTIzDGgIyALCk0BTI7OgSOiP2WU0CGGIiBkac6gGGcNYQxAVVYiksoa6t34VJOjU43OUD1nNHHNBCpB6d1FwV2OFkydn0ao6D9Yg4t"
	//SSOM
	var jwtSecretKey = "Oa4OkXlw98ECVxikEwVlTmiBybXuv4zLqbHXSoO59o5KXMdh05ygzqQ8CLIfMWPEgF3NTB6MRkLT0atn2JBklGD0rNaWN3lmLRglDqohRsPVFcRt2tSr48ark88nNxlmTr1b5nuqMGEcwYFrQraUPeW1UvrvqQton0YeOF5bT2ss9Gpyv2Yc9Jp0t4Rkzq9ZfaQKR77r472pZcW8MpnQ0ZV5UmfFlSyjFpn4YI9a0Ky6mb3SzlFesytXWcgByObM";
	var jwtTokenExpire = 1000000
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Second * time.Duration(jwtTokenExpire)).Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		log.Println("Error while signing the token for online support")
		return "", err
	} else {
		return tokenString, nil
	}
}

func ValidateJwtTokenForOnlineSupport(tockenString string) error{
	//var jwtSecretKey = "QJZundppmWpyAsqpolsIqZ2KcDC0cZBx2ZVcWcu8dgmMKx1hToHAhoerPRan30fMYxcKLAoMpbW3QC4v7MOYAeQH0o4KzZxmQJDNfFdzBCfFVTkRcwcHXxyoqgkyR3I6ZE0ZTAdT9Q9l1d0We8oJRTFZhFcTIzDGgIyALCk0BTI7OgSOiP2WU0CGGIiBkac6gGGcNYQxAVVYiksoa6t34VJOjU43OUD1nNHHNBCpB6d1FwV2OFkydn0ao6D9Yg4t"
	//SSOM
	var jwtSecretKey = "Oa4OkXlw98ECVxikEwVlTmiBybXuv4zLqbHXSoO59o5KXMdh05ygzqQ8CLIfMWPEgF3NTB6MRkLT0atn2JBklGD0rNaWN3lmLRglDqohRsPVFcRt2tSr48ark88nNxlmTr1b5nuqMGEcwYFrQraUPeW1UvrvqQton0YeOF5bT2ss9Gpyv2Yc9Jp0t4Rkzq9ZfaQKR77r472pZcW8MpnQ0ZV5UmfFlSyjFpn4YI9a0Ky6mb3SzlFesytXWcgByObM";
	_, err := jwt.ParseWithClaims(tockenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})
	if err != nil{
		return err
	}

	return nil
}