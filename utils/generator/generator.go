package generator

import (
	"math/rand"
	"time"
	"unsafe"
)

func GenerateRandomChar(size int) string {
	rand.Seed(time.Now().Add(time.Duration(size) * time.Second).Unix())

	var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]byte, size)

	rand.Read(b)

	for i := 0; i < size; i++ {
		b[i] = alphabet[b[i]%byte(len(alphabet))]
	}
	return *(*string)(unsafe.Pointer(&b))
}

