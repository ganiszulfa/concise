package generate

import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func RandAlphabetsLowerCase(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
