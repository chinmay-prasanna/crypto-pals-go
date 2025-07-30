package test1

import (
	"encoding/hex"
	"fmt"
	"log"
)

func Xor() {
	hexstr1, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b1 := hexstr1
	hexstr2, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	b2 := hexstr2
	if len(b1) != len(b2) {
		log.Fatal("Buffer length should be equal")
	}
	xorstr := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		x := b1[i] ^ b2[i]
		xorstr[i] = x
	}
	fmt.Print(hex.EncodeToString(xorstr))
}
