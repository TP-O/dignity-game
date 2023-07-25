package pkg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func SignWithHMAC(key string, msg string) (signature string, err error) {
	hash := hmac.New(sha256.New, []byte(key))
	if _, err = hash.Write([]byte(msg)); err != nil {
		return
	}

	signature = hex.EncodeToString(hash.Sum(nil))
	return
}
