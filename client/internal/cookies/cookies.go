package cookies

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type claim struct {
	jwt.RegisteredClaims
	UserID uint32
}

func NewCookie(w http.ResponseWriter, userID uint32, secretKey string) (string, error) {

	token, err := createToken(userID, secretKey)
	if err != nil {
		return "", err
	}

	cookie := &http.Cookie{
		Name:  "Token",
		Value: token,
		Path:  "/",
	}

	http.SetCookie(w, cookie)

	return token, nil
}

func createToken(userID uint32, secretKey string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		claim{
			RegisteredClaims: jwt.RegisteredClaims{},
			UserID:           userID,
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Validation(token, secretkey string) bool {
	tok, err := jwt.Parse(token,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(secretkey), nil
		})

	if err != nil || !tok.Valid {
		return false
	}
	return true
}
