package types

import (
	"string_encryption/encrypt"
)

type File struct {
	Name string
	Code []byte
	Path string
}

type String struct{
	String string
	Encrypted string
}

func (s String) Encrypt() string {
	return encrypt.Encrypt(s.String)
}