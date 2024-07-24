package token

import (
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	common "github.com/saas0503/factory-common"
	"time"
)

type Detail struct {
	TokenID   string
	Token     *string
	UserID    string
	ExpiresIn *int64
}

func Generate(userID string, ttl time.Duration, privateKey string) (*Detail, error) {
	now := time.Now().UTC()
	td := &Detail{
		ExpiresIn: new(int64),
		Token:     new(string),
	}
	*td.ExpiresIn = now.Add(ttl).Unix()
	td.TokenID = common.UUID()
	td.UserID = userID

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode token private key: %w", err)
	}

	key, er := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if er != nil {
		return nil, fmt.Errorf("could not parse token private key: %w", er)
	}

	atClaims := make(jwt.MapClaims)
	atClaims["sub"] = userID
	atClaims["id"] = td.TokenID
	atClaims["exp"] = td.ExpiresIn
	atClaims["iat"] = now.Unix()

	*td.Token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims).SignedString(key)
	if err != nil {
		return nil, fmt.Errorf("could not sign token: %w", err)
	}

	return td, nil
}

func Verify(token string, publicKey string) (*Detail, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode token public key: %w", err)
	}

	key, er := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if er != nil {
		return nil, fmt.Errorf("could not parse token public key: %w", err)
	}

	parsedToken, errParse := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if errParse != nil {
		return nil, fmt.Errorf("could not parse token: %w", errParse)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &Detail{
		TokenID: fmt.Sprint(claims["id"]),
		UserID:  fmt.Sprint(claims["sub"]),
	}, nil
}
