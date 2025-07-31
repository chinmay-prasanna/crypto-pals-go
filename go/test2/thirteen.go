package test2

import (
	"fmt"
	"strings"
)

var key []byte

func init() {
	key = RandomKey(16)
}

func ProfileFor(email string) (string, error) {
	email = strings.ReplaceAll(email, "&", "")
	email = strings.ReplaceAll(email, "=", "")
	str := fmt.Sprintf("email=%s&uid=10&role=user", email)
	return str, nil
}

func ParseProfile(data string) map[string]string {
	parts := strings.Split(data, "&")
	result := make(map[string]string)
	for _, part := range parts {
		kv := strings.Split(part, "=")
		if len(kv) == 2 {
			result[kv[0]] = kv[1]
		}
	}
	return result
}

func DecryptProfile(cipher []byte) map[string]string {
	plain := DecryptECB(cipher, key)
	return ParseProfile(string(plain))
}

func EncryptProfile(email string) []byte {
	profile, _ := ProfileFor(email)
	enc := EncryptECB([]byte(profile), key)
	return enc
}

func PerformAttack() {
	blockSize := 16
	pLen := blockSize - len("email=")
	prefix := strings.Repeat("A", pLen)
	admin := "admin"
	padmin := PKCS7Padding([]byte(admin), blockSize)
	email := prefix + string(padmin)
	fmt.Print([]byte(email))
	c1 := EncryptProfile(email)
	aBlock := c1[blockSize : 2*blockSize]
	email2 := strings.Repeat("A", 13)
	normal := EncryptProfile(email2)
	forged := append(normal[:2*blockSize], aBlock...)

	result := DecryptProfile(forged)
	fmt.Println((result))
}
