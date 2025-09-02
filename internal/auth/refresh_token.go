package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() (string, error) {
	randString := make([]byte, 32)
	rand.Read(randString)
	return hex.EncodeToString(randString), nil
}
