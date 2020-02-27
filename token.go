package jwt

import (
	"errors"
	"strconv"
	"strings"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

var _signingKey []byte
var _expireSecond time.Duration

// Default Setup
func init() {
	Setup("123456", 60)
}

// Setup : privatte key , and expire second
func Setup(key string, t time.Duration) {
	_signingKey = []byte(key)
	_expireSecond = t
}

// Build ...
func Build(in map[string]string) (token string, err error) {
	exp := timeToString(time.Now().Add(_expireSecond))
	claims := jwtgo.MapClaims{"exp": exp}
	for k, v := range in {
		claims[k] = v
	}

	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	return token.SignedString(_signingKey)
}

// Verify ...
func Verify(tokenString string) (result map[string]string, err error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwtgo.Parse(tokenString, func(token *jwtgo.Token) (interface{}, error) {
		return _signingKey, nil
	})
	if err != nil {
		return
	}

	claims := token.Claims.(jwtgo.MapClaims)
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
func timeToString(t time.Time) (timeStr string) {
	v, _ := strconv.ParseInt(
		strconv.FormatInt(t.UTC().UnixNano(), 10)[:13],
		10,
		64,
	)
	return strconv.Itoa(int(v))
}
