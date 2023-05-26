package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value, salt string) string {
	m := md5.New()
	m.Write([]byte(salt))
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

// sha256 encryption
func EncodeSHA256(value, salt string) string {
	h := sha256.New()
	h.Write([]byte(salt))
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}
