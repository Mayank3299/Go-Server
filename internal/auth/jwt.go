package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string

const (
	// TokenTypeAccess -
	TokenTypeAccess TokenType = "chirpy-access"
)

func MakeJWT(userID uuid.UUID, tokenSecret string) (string, error) {
	signingMethod := jwt.SigningMethodHS256
	claims := jwt.RegisteredClaims{
		Issuer:    string(TokenTypeAccess),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
		Subject:   userID.String(),
	}

	token := jwt.NewWithClaims(signingMethod, claims)
	res, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		log.Printf("Couldn't create jwt: %v", err)
		return "", err
	}

	return res, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(tokenSecret), nil
		},
	)

	if err != nil {
		log.Printf("Invalid or Expired token: %v", err)
		return uuid.Nil, err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.Nil, err
	}

	if issuer != string(TokenTypeAccess) {
		return uuid.Nil, errors.New("invalid issuer")
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return uuid.Parse(claims.Subject)
	}

	return uuid.Nil, fmt.Errorf("invalid token claims")
}

func GetBearerToken(headers http.Header) (string, error) {
	authorizationHeader := headers.Get("Authorization")
	if authorizationHeader == "" {
		return "", fmt.Errorf("please provide authorization token")
	}

	parts := strings.Split(authorizationHeader, " ")
	if parts[0] != "Bearer" || len(parts) != 2 {
		return "", errors.New("invalid authorization header")
	}

	tokenString := parts[len(parts)-1]
	return tokenString, nil
}

func GetAPIKey(headers http.Header) (string, error) {
	authorizationHeader := headers.Get("Authorization")
	if authorizationHeader == "" {
		return "", fmt.Errorf("please provide authorization token")
	}

	parts := strings.Split(authorizationHeader, " ")
	if parts[0] != "ApiKey" || len(parts) != 2 {
		return "", errors.New("invalid authorization header")
	}

	tokenString := parts[len(parts)-1]
	return tokenString, nil
}
