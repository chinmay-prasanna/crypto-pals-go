package test1

import (
	"unicode"
)

func ScoreEnglish(text []byte) float64 {
	freq := map[byte]float64{
		'a': 8.12, 'b': 1.49, 'c': 2.71, 'd': 4.32, 'e': 12.0, 'f': 2.30,
		'g': 2.03, 'h': 5.92, 'i': 7.31, 'j': 0.10, 'k': 0.69, 'l': 3.98,
		'm': 2.61, 'n': 6.95, 'o': 7.68, 'p': 1.82, 'q': 0.11, 'r': 6.02,
		's': 6.28, 't': 9.10, 'u': 2.88, 'v': 1.11, 'w': 2.09, 'x': 0.17,
		'y': 2.11, 'z': 0.07, ' ': 13.0,
	}
	score := 0.0
	for _, b := range text {
		b = byte(unicode.ToLower(rune(b)))
		if val, ok := freq[b]; ok {
			score += val
		} else {
			score -= 1.5
		}
	}

	return score
}

func GetEncodingKey(bstr []byte) (byte, float64, []byte) {
	// bstr, _ := hex.DecodeString(str)
	bestScore := -1.0
	var key byte
	for i := 0; i < 256; i++ {
		decoded := make([]byte, len(bstr))
		for j := 0; j < len(bstr); j++ {
			decoded[j] = bstr[j] ^ byte(i)
		}
		score := ScoreEnglish(decoded)
		if score > bestScore {
			bestScore = score
			key = byte(i)
		}
	}
	decoded := make([]byte, len(bstr))
	for i := 0; i < len(bstr); i++ {
		decoded[i] = bstr[i] ^ key
	}
	return key, bestScore, decoded

}
