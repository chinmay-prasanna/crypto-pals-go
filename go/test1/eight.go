package test1

import "fmt"

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

func ScanForECBEncryption(str []byte) bool {
	blockSize := 16
	if len(str)%blockSize != 0 {
		return false
	}
	c := CountDuplicates(str)
	fmt.Print(c)
	return c > 1
}
