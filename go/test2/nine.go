package test2

import (
	"bytes"
)

func PKCS7Padding(data []byte, bs int) []byte {
	padding := bs - (len(data) % bs)
	pad := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, pad...)
}
