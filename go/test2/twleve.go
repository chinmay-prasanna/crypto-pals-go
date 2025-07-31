package test2

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"kew/cryptopal/test1"
)

func oracle(input []byte) []byte {
	str := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`
	unknown, _ := base64.StdEncoding.DecodeString(str)
	plaintext := append(input, unknown...)
	ciphertext := EncryptECB(plaintext, key)
	return ciphertext
}

func GetECBKey() {
	size := 0
	lastLen := len(oracle([]byte{}))
	for i := 0; i < 64; i++ {
		input := bytes.Repeat([]byte{byte(i)}, i)
		cipher := oracle(input)
		newLen := len(cipher)
		if newLen > lastLen {
			size = newLen - lastLen
			break
		}
	}
	repeatInput := bytes.Repeat([]byte("A"), size*4)
	ciphertext := oracle(repeatInput)
	if test1.ScanForECBEncryption(ciphertext) {
		fmt.Println("ECB mode detected")
	} else {
		fmt.Println("ECB mode NOT detected")
	}
	recovered := []byte{}
	for n := 1; n < size; n++ {
		input := bytes.Repeat([]byte("A"), (size - n))
		var gByte byte
		bMap := map[string]byte{}
		for i := 0; i < 256; i++ {
			guess := append(input, recovered...)
			guess = append(guess, byte(i))
			gByte = byte(i)
			cipher := oracle(guess)
			bMap[(string(cipher[:size]))] = gByte
		}
		output := oracle(input)[:size]
		outStr := string(output)
		for k, v := range bMap {
			if k == outStr {
				recovered = append(recovered, v)
				gByte = v
				break
			}
		}
	}
	fmt.Print(string(recovered))
}
