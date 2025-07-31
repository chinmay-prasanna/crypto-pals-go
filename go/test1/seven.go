package test1

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
)

func ECBDecryption() {
	data, _ := os.ReadFile("file.md")
	bstr, _ := base64.StdEncoding.DecodeString(string(data))
	key := []byte("YELLOW SUBMARINE")

	block, _ := aes.NewCipher(key)
	blockBuf := make([]byte, aes.BlockSize)
	decoded := make([]byte, len(bstr))
	for i := 0; i < len(bstr); i += aes.BlockSize {
		encoded := bstr[i : i+aes.BlockSize]
		block.Decrypt(blockBuf, encoded)
		decoded = append(decoded, blockBuf...)
	}

	fmt.Print(string(decoded))
}
