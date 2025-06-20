package encrypt

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func Encrypt (str string) string {
	key := make([]byte, 4)
	encStr := []byte(str)

	_, err := rand.Read(key)
	if err != nil {
		log.Fatal("[-] Failed while generating the key :", err.Error())
	}

	for i, b := range encStr {
		encStr[i] = b ^ key[i%len(key)]
	}

	res := hex.EncodeToString(key) + hex.EncodeToString(encStr)
	return res
}