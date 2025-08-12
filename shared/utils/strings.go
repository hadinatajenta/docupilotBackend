package utils

import (
	crand "crypto/rand"
	"math/big"
)

const idAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandID generates random string with max length 8 (secure).
func RandID(n int) (string, error) {
	if n > 8 {
		n = 8
	}
	if n <= 0 {
		n = 1
	}

	out := make([]byte, n)
	max := big.NewInt(int64(len(idAlphabet)))
	for i := 0; i < n; i++ {
		idx, err := crand.Int(crand.Reader, max)
		if err != nil {
			return "", err
		}
		out[i] = idAlphabet[idx.Int64()]
	}
	return string(out), nil
}
