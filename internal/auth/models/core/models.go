package core

import (
	"encoding/base64"
	"fmt"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
)

type Auth struct {
	Id          string `db:"uuid"`
	Email       string `db:"email"`
	IsConfirmed bool   `db:"is_confirmed"`
	UserPassword
}

type UserPassword struct {
	Hash string `db:"password_hash"`
	Salt string `db:"password_salt"`
}

// Init generates salt and hash with given password and fills corresponding fields.
func (up *UserPassword) Init(password string) error {
	salt, err := utils.GetSalt()
	if err != nil {
		return fmt.Errorf("error generating salt: %s", err)
	}
	hash, err := utils.GetHash512(password, salt)
	if err != nil {
		return fmt.Errorf("error generating hash: %s", err)
	}

	up.Salt = base64.URLEncoding.EncodeToString(salt)
	up.Hash = base64.URLEncoding.EncodeToString(hash)

	return nil
}

// Validate checks if the given password is the one that is stored.
func (up *UserPassword) Validate(password string) error {
	salt, err := base64.URLEncoding.DecodeString(up.Salt)
	if err != nil {
		return fmt.Errorf("error decoding user's salt: %s", err)
	}

	hash, err := utils.GetHash512(password, salt)
	if err != nil {
		return fmt.Errorf("error generating hash: %s", err)
	}

	if base64.URLEncoding.EncodeToString(hash) != up.Hash {
		return constants.ErrPasswordMismatch
	}

	return nil
}
