package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const JWT_SECRET = "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5"

type Content struct {
	Id       int
	Username string
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}
func VerifyToken(tokenStr string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenStr, secretFunc(JWT_SECRET))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("get token content error")
	}
	return claims, err
}
func Sign(c *Content) (tokenStr string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":       c.Id,
		"Username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenStr, err = t.SignedString([]byte(JWT_SECRET))
	return
}
