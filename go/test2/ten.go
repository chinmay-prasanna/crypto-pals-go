package test2

import (
	"crypto/aes"
)

func XOR(a, b []byte) []byte {
	xored := []byte{}
	for i := 0; i < len(a); i++ {
		xored = append(xored, a[i]^b[i])
	}
	return xored
}

func EncryptCBC(bstr []byte, key []byte, iv []byte) []byte {
	bstr = PKCS7Padding(bstr, aes.BlockSize)
	block, _ := aes.NewCipher(key)
	encoded := []byte{}
	prev := iv
	for i := 0; i < len(bstr); i += aes.BlockSize {
		encrypted := make([]byte, aes.BlockSize)
		chunk := bstr[i : i+aes.BlockSize]
		eChunk := XOR(chunk, prev)
		block.Encrypt(encrypted, eChunk)
		encoded = append(encoded, encrypted...)
		prev = encrypted
	}

	return encoded
}

func DecryptCBC(bstr []byte, key []byte, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	decoded := []byte{}
	prev := iv
	for i := 0; i < len(bstr); i += aes.BlockSize {
		decrypted := make([]byte, aes.BlockSize)
		chunk := bstr[i : i+aes.BlockSize]
		block.Decrypt(decrypted, chunk)
		dChunk := XOR(decrypted, prev)
		decoded = append(decoded, dChunk...)
		prev = chunk
	}

	return PKCS7Padding(decoded, aes.BlockSize)
}

func EncryptECB(bstr []byte, key []byte) []byte {
	bstr = PKCS7Padding(bstr, aes.BlockSize)
	block, _ := aes.NewCipher(key)
	encoded := []byte{}
	for i := 0; i < len(bstr); i += aes.BlockSize {
		encrypted := make([]byte, aes.BlockSize)
		chunk := bstr[i : i+aes.BlockSize]
		block.Encrypt(encrypted, chunk)
		encoded = append(encoded, encrypted...)
	}
	return encoded
}

func DecryptECB(bstr []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	decoded := []byte{}
	for i := 0; i < len(bstr); i += aes.BlockSize {
		decrypted := make([]byte, aes.BlockSize)
		chunk := bstr[i : i+aes.BlockSize]
		block.Decrypt(decrypted, chunk)
		decoded = append(decoded, decrypted...)
	}

	return PKCS7Padding(decoded, aes.BlockSize)
}
