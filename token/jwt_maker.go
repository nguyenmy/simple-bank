package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var ErrorInvalidToken = errors.New("invalid token")

type JWTMaker struct {
	secretKey string
}

func NewJwtMaker(secretKey string) (*JWTMaker, error) {
	if len(secretKey) < 32 {
		return nil, errors.New("secret key len must larger then 32")
	}

	return &JWTMaker{
		secretKey: secretKey,
	}, nil
}

func (j *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {

	payload := NewPayload(username, duration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTMaker) Verify(tokenS string) (*Payload, error) {
	var payload Payload
	token, err := jwt.ParseWithClaims(tokenS, &payload, func(t *jwt.Token) (interface{}, error) {
		// check algorithm
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(j.secretKey), nil
	})

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrorExpiredToken) {
			return nil, ErrorExpiredToken
		}
		return nil, ErrorInvalidToken
	}

	if !token.Valid {
		return nil, ErrorInvalidToken
	}

	return &payload, nil
}
