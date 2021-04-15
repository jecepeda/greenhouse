package crypt

import (
	"github.com/jecepeda/greenhouse/server/gerror"
	"golang.org/x/crypto/bcrypt"
)

// BEncrypter is a encrypter that uses the bcrypt algorithm
// to hash passwords.
type BEncrypter struct{}

// EncryptPassword takes a plain password and hashes it using a cryptographic salt
// function
func (b BEncrypter) EncryptPassword(s string) ([]byte, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, gerror.Wrap(err, "encrypt password")
	}
	return hashedBytes, nil
}

// CheckPassword compares a plain text password with the hashed one
// Returns nil if success, error if not
func (b BEncrypter) CheckPassword(existing []byte, new string) error {
	incoming := []byte(new)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
