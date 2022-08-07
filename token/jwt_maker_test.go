package token

import (
	"fmt"
	"go-simple-bank/db/util"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvalidSecretSign(t *testing.T) {
	_, err := NewJwtMaker("abcd2")
	assert.Error(t, err)
}

func createToken(j *JWTMaker) (string, error) {

	token, err := j.CreateToken("mynguyen", time.Minute*5)
	return token, err
}
func TestExpiredToken(t *testing.T) {
	jmaker, err := NewJwtMaker(util.RandString(32))
	assert.NoError(t, err)

	token, err := jmaker.CreateToken("mynguyen", -time.Minute*5)
	fmt.Println(token)
	assert.NoError(t, err)
	_, err = jmaker.Verify(token)
	assert.Error(t, err)
	require.EqualError(t, err, ErrorExpiredToken.Error())
}

func TestVerifyToken(t *testing.T) {
	jmaker, err := NewJwtMaker(util.RandString(32))
	assert.NoError(t, err)

	token, err := createToken(jmaker)
	fmt.Println(token)
	assert.NoError(t, err)
	p, err := jmaker.Verify(token)
	assert.NoError(t, err)
	assert.Equal(t, "mynguyen", p.Username)
}

func TestInvalidTokenAlg(t *testing.T) {
	payload := NewPayload(util.RandomOwner(), time.Second*4)
	token := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	tokenS, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	assert.NoError(t, err)

	jmaker, err := NewJwtMaker(util.RandString(32))
	assert.NoError(t, err)

	_, err = jmaker.Verify(tokenS)
	assert.Error(t, err)
	require.EqualError(t, err, ErrorInvalidToken.Error())

}
