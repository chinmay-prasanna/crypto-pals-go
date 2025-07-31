package test2

import (
	"bytes"
	"math/rand"
)

func encrypt(bstr []byte, key []byte) []byte {
	lp := rand.Intn(10)
	rp := rand.Intn(10)
	lb := bytes.Repeat([]byte{byte(lp)}, lp)
	rb := bytes.Repeat([]byte{byte(rp)}, rp)
	bstr = append(lb, bstr...)
	bstr = append(bstr, rb...)

	choice := rand.Intn(2)
	if choice == 0 {
		return EncryptECB(bstr, key)
	}
	iv := RandomKey(16)
	return EncryptCBC(bstr, key, iv)
}

func RandomKey(kl int) []byte {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	bstr := []byte{}
	for i := 0; i < kl; i++ {
		bstr = append(bstr, byte(chars[rand.Intn(len(chars))]))
	}
	return bstr
}
