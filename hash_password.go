package main

import (
	"crypto/hmac"
	"crypto/sha256"
)

type HashPassword struct {
	Password   string
	PassSecret []byte
}

func (h HashPassword) Validate() (validations []ErrValidation) {

	if h.Password == "" {
		validations = append(validations, ErrValidation{"password", "password is empty"})
	}
	if len(h.Password) < 8 {
		validations = append(validations, ErrValidation{"password", "password is short"})
	}

	if h.PassSecret == nil {
		validations = append(validations, ErrValidation{"pass_secret", "pass secret is empty"})
	}

	if len(validations) > 0 {
		return validations
	}
	return
}

func (h *HashPassword) HashPassword() []byte {
	hash := hmac.New(sha256.New, h.PassSecret)
	hash.Write([]byte(h.Password))
	return hash.Sum(nil)
}

func (h *HashPassword) CheckPassword(hashedPassword []byte) bool {
	return hmac.Equal(hashedPassword, h.HashPassword())
}
