package test1

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"sort"
)

func GetHammingDistance(b1 []byte, b2 []byte) float64 {
	dist := 0.0
	for i := 0; i < len(b1); i++ {
		xor := b1[i] ^ b2[i]
		for xor > 0 {
			dist++
			xor &= xor - 1
		}
	}
	return dist
}

func BreakRepatingKey() {
	data, _ := os.Open("file.md")
	scanner := bufio.NewScanner(data)
	encStr := ""
	for scanner.Scan() {
		line := scanner.Text()
		encStr += line
	}
	bstr, _ := base64.StdEncoding.DecodeString(encStr)
	kmap := map[int]float64{}
	for k := 2; k < 40; k++ {
		nb := 4
		td := 0.0
		pairs := 0
		for n := 0; n < nb; n++ {
			s1 := n * k
			s2 := (n + 1) * k
			if s2 > len(bstr) {
				break
			}
			c1 := bstr[s1 : s1+k]
			c2 := bstr[s2 : s2+k]
			dist := GetHammingDistance(c1, c2)
			td += dist
			pairs++
		}
		if pairs > 0 {
			avg := td / float64(pairs)
			navg := avg / float64(k)
			kmap[k] = navg
		}
	}
	type KeyValue struct {
		Key   int
		Value float64
	}
	ordered := []KeyValue{}
	for k, v := range kmap {
		ordered = append(ordered, KeyValue{k, v})
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Value < ordered[j].Value
	})
	keys := ordered[:3]
	for _, k := range keys {
		blocks := [][]byte{}
		for i := 0; i < len(bstr); i += k.Key {
			end := i + k.Key
			if end > len(bstr) {
				end = len(bstr)
			}
			chunk := bstr[i:end]
			blocks = append(blocks, chunk)
		}
		transposed := [][]byte{}
		for i := 0; i < k.Key; i++ {
			var col []byte
			for _, j := range blocks {
				if i < len(j) {
					col = append(col, j[i])
				}
			}
			transposed = append(transposed, col)
		}
		var keystr []byte
		for _, col := range transposed {
			maxScore := -1.0
			var bestByte byte
			for b := 0; b < 256; b++ {
				xored := make([]byte, len(col))
				for i := range col {
					xored[i] = col[i] ^ byte(b)
				}
				score := ScoreEnglish(xored)
				if score > maxScore {
					maxScore = score
					bestByte = byte(b)
				}
			}
			keystr = append(keystr, bestByte)
		}
		plaintext := make([]byte, len(bstr))
		for i := 0; i < len(bstr); i++ {
			plaintext[i] = bstr[i] ^ keystr[i%len(keystr)]
		}

		fmt.Printf("\nKeysize: %d | Key: %s\n", k.Key, string(keystr))
		fmt.Printf("Decrypted output (first 300 chars):\n%s\n", string(plaintext[:300]))
	}
}
