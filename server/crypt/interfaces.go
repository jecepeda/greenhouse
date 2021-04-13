package crypt

type Encrypter interface {
	EncryptPassword(s string) ([]byte, error)
	CheckPassword(existing []byte, new string) error
}
