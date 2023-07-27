package Authorization

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/golang-jwt/jwt"

	"upbit/Logger"

	"github.com/google/uuid"
)

func (*Authorization) GetAuthoriztionToken(queryString string) string {
	var AuthorizationToken string

	claims := jwt.MapClaims{}
	claims["access_key"] = "UNHxQaK9kRTic4v8tUT9AEzZs4bPVduWY6zEqRNZ"
	claims["nonce"] = uuid.New().String()

	if queryString != "" {
		hash := sha512.Sum512([]byte(queryString))
		hashSlice := hash[:] // [64]byte -> []byte (convert static byte to slice)
		hexHash := hex.EncodeToString(hashSlice)

		claims["query_hash"] = hexHash
		claims["query_hash_alg"] = "SHA512"
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := at.SignedString([]byte("5uJnZ0qujgQxLSDnmqXTWy7jRr1vxpiL8Joq2rjY"))
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return AuthorizationToken
	}
	AuthorizationToken = "Bearer " + token

	return AuthorizationToken
}
