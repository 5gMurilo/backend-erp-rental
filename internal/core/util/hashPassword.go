package util

import (
	"crypto/sha256"
)

func HashPassword(password string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		return "", err
	}
	return string(h.Sum(nil)), err
}
