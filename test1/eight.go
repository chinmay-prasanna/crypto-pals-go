package test1

import (
	"bufio"
	"encoding/hex"
	"os"
)

func CountDuplicates(str []byte) int {
	blocks := map[string]int{}

	for i := 0; i < len(str); i += 16 {
		if i+16 > len(str) {
			break
		}
		block := str[i : i+16]
		blocks[string(block)]++
	}
	count := 0
	for _, v := range blocks {
		if v > 1 {
			count += v - 1
		}
	}
	return count
}

func ScanForECBEncryption() {
	data, _ := os.Open("file.md")
	scanner := bufio.NewScanner(data)
	hexstr := [][]byte{}
	for scanner.Scan() {
		hs, _ := hex.DecodeString(scanner.Text())
		hexstr = append(hexstr, hs)
	}
	blockSize := 16
	dupes := [][]byte{}
	for _, str := range hexstr {
		if len(str)%blockSize != 0 {
			continue
		}
		c := CountDuplicates(str)
		if c > 1 {
			dupes = append(dupes, str)
		}
	}
}
