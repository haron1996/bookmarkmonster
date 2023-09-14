package token

import (
	"errors"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func VerifyToken(signedToken string) (*PayLoad, error) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		return nil, err
	}

	publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(config.PublicKeyHex)
	if err != nil {
		return nil, err
	}

	parser := paseto.NewParser()

	token, err := parser.ParseV4Public(publicKey, signedToken, nil)
	if err != nil {
		return nil, err
	}

	var payload PayLoad

	if err := token.Get("payload", &payload); err != nil {
		return nil, err
	}

	if time.Now().UTC().After(payload.Expiry) {
		return nil, errors.New("token is expired")
	}

	return &payload, nil
}
