package aes

import (
	"crypto/sha1"
	"encoding/base64"
)

func createHash(key string) string {
	hasher := sha1.New()
	hasher.Write([]byte(key))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
