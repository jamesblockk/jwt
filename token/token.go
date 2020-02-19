package token

import (
	"errors"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var signingKey []byte
var validHour time.Duration

// Setup : privatte key , and valid hour  
func Setup(key string, t time.Duration) {
	signingKey = []byte(key)
	validHour = t
}

// Build ...
func Build(in map[string]string) (string, error) {
	ad := time.Duration(time.Hour * validHour)
	exp := timeToString(time.Now().Add(ad))
	claims := jwt.MapClaims{"exp": exp}
	for k, v := range in {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

// Verify ...
func Verify(tokenString string) (result map[string]string, err error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	exp, ok := claims["exp"].(string)
	if !ok {
		err = errors.New("parse fail")
		return
	}

	if exp < timeToString(time.Now()) {
		err = errors.New("exp fail : out of valid range")
		return
	}

	result = map[string]string{}
	for k := range claims {
		v, ok := claims[k].(string)
		if !ok {
			err = errors.New("parse fail")
			return
		}
		result[k] = v
	}

	return result, nil
}

// -- private function ------------

// Default Setup
func init() {
	Setup("123456", 1)
}

func timeToString(t time.Time) string {
	str := strconv.FormatInt(t.UTC().UnixNano(), 10)
	v, _ := strconv.ParseInt(str[:13], 10, 64)
	return strconv.Itoa(int(v))
}
