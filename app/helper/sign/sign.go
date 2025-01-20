package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func String(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
