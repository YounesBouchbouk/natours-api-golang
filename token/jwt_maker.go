package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {

	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey: secretKey}, nil

}

func (maker *JWTMaker) CreateToken(email string, role string, duration time.Duration) (string, *Payload, error) {

	payload, err := NewPayload(email, role, duration)

	if err != nil {
		return "", nil, fmt.Errorf("error in creating token payload: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":        payload.ID,
		"Email":     payload.Email,
		"Role":      payload.Role,
		"IssuedAt":  payload.IssuedAt,
		"ExpiredAt": payload.ExpiredAt,
	})

	tokenString, err := token.SignedString([]byte(maker.secretKey))

	if err != nil {
		return "", nil, fmt.Errorf("error in creating token: %w", err)

	}

	return tokenString, payload, nil
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, ErrInvalidToken
	}

	fmt.Print(payload)

	return payload, nil

}
