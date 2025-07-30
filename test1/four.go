package test1

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func GetKeyInFile() {
	data, _ := os.Open("file.md")
	scanner := bufio.NewScanner(data)
	globalBestScore := 0.0
	cipherText := []byte{}
	var globalKey byte
	for scanner.Scan() {
		bstr, _ := hex.DecodeString(scanner.Text())
		bestScore := -1.0
		var key byte
		var ctext []byte
		for i := 0; i < 256; i++ {
			decoded := make([]byte, len(bstr))
			for j := 0; j < len(bstr); j++ {
				decoded[j] = bstr[j] ^ byte(i)
			}
			score := ScoreEnglish(decoded)
			if score > bestScore {
				bestScore = score
				key = byte(i)
				ctext = decoded
			}
		}
		if bestScore > globalBestScore {
			globalBestScore = bestScore
			globalKey = key
			cipherText = ctext
		}
	}
	fmt.Println(globalKey, "-", string(cipherText))
}
