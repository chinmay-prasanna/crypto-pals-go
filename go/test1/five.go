package test1

import (
	"encoding/hex"
	"fmt"
)

func EncryptRepeatingXor() {
	bstr := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
	kstr := []byte("ICE")
	encrypted := []byte{}
	for i := 0; i < len(bstr); i++ {
		encrypted = append(encrypted, bstr[i]^kstr[i%len(kstr)])
	}
	fmt.Println(hex.EncodeToString(encrypted))
}
