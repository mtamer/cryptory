package cryptory

import (
	"errors"
	"strings"
)

// Ceasar's Encryption Cipher
// Developed by Julius Ceasar
func CeasarEnc(msg string) (string, error) {
	upper := strings.ToUpper(msg)
	var encrypted []rune
	for _, value := range upper {
		value += 3
		if value > 90 {
			value -= 26
		}
		if value < 65 || value > 90 {
			return "", errors.New("Not an alphabetic string.")
		}
		encrypted = append(encrypted, value)
	}
	return string(encrypted), nil
}

// Ceasar's Decryption Cipher
func CeasarDec(encrypted string) (string, error) {
	upper := strings.ToUpper(encrypted)
	var msg []rune
	for _, value := range upper {
		value -= 3
		if value < 65 {
			value += 26
		}
		if value < 65 || value > 90 {
			return "", errors.New("Not an alphabetic string.")
		}
		msg = append(msg, value)
	}
	return string(msg), nil
}

// Mono Alphabetic Encryption Cipher
func MonoAlphaEnc(msg string) (string, error) {
	upper := strings.ToUpper(msg)
	var encrypted []rune
	for _, value := range upper {
		if value < 65 || value > 90 {
			return "", errors.New("Not an alphabetic string.")
		}
		encrypted = append(encrypted, monoMap[value])
	}
	return string(encrypted), nil
}

// Mono Alphabetic Decryption Cipher
func MonoAlphaDec(encrypted string) (string, error) {
	upper := strings.ToUpper(encrypted)
	var msg []rune
	for _, value := range upper {
		if value < 65 || value > 90 {
			return "", errors.New("Not an alphabetic string.")
		}
		msg = append(msg, monoMapRev[value])
	}
	return string(msg), nil
}

// Character Map for MonoAlphabetic Encryption
var monoMap = map[rune]rune{'A': 'Z', 'B': 'W', 'C': 'X', 'D': 'Y', 'E': 'O', 'F': 'P', 'G': 'R', 'H': 'Q', 'I': 'A', 'J': 'H', 'K': 'C', 'L': 'B', 'M': 'E', 'N': 'S', 'O': 'U', 'P': 'T', 'Q': 'V', 'R': 'F', 'S': 'G', 'T': 'J', 'U': 'L', 'V': 'K', 'W': 'M', 'X': 'N', 'Y': 'D', 'Z': 'I'}

// Character Map for MonoAlphabetic Decryption
var monoMapRev = map[rune]rune{'Z': 'A', 'W': 'B', 'X': 'C', 'Y': 'D', 'O': 'E', 'P': 'F', 'R': 'G', 'Q': 'H', 'A': 'I', 'H': 'J', 'C': 'K', 'B': 'L', 'E': 'M', 'S': 'N', 'U': 'O', 'T': 'P', 'V': 'Q', 'F': 'R', 'G': 'S', 'J': 'T', 'L': 'U', 'K': 'V', 'M': 'W', 'N': 'X', 'D': 'Y', 'I': 'Z'}

// Hill Encryption Cipher
// Developed by Lester Hill
func HillEnc(msg string) (string, string, error) {
	if len(msg)%2 != 0 {
		return "", "", errors.New("Message must be an even length!")
	}
	key := genKey()
	result := make([]byte, len(msg))

	sets := len(msg) / 2
	for i := 0; i < sets; i++ {
		offset := i * 2
		result[offset+0] = (((msg[offset+0]-65)*key[0])+((msg[offset+1]-65)*key[1]))%26 + 65
		result[offset+1] = (((msg[offset+0]-65)*key[2])+((msg[offset+1]-65)*key[3]))%26 + 65
	}

	return string(result), string(key), nil
}

// Hill Decryption Cipher
func HillDec(enc, key string) (string, error) {
	if len(enc)%2 != 0 {
		return "", errors.New("Message must be an even length!")
	}
	dKey, err := genDKey(key)
	if err != nil {
		return "", err
	}
	result := make([]byte, len(enc))

	sets := len(enc) / 2
	for i := 0; i < sets; i++ {
		offset := i * 2
		result[offset+0] = byte((int(enc[offset+0]-65)*int(dKey[0])+int(enc[offset+1]-65)*int(dKey[1]))%26 + 65)
		result[offset+1] = byte((int(enc[offset+0]-65)*int(dKey[2])+int(enc[offset+1]-65)*int(dKey[3]))%26 + 65)
	}

	return string(result), nil
}

// Generate key
func genKey() []byte {
	return []byte{3, 3, 2, 5}
}

func genDKey(key string) ([]byte, error) {
	dkey := []byte{key[3], 26 - key[1], 26 - key[2], key[0]}
	det := (key[3] * key[0]) - (key[2] * key[1])
	i := uint8(1)
	for (det*i)%26 != 1 {
		i++
	}
	for j, _ := range dkey {
		dkey[j] = ((dkey[j] * i) % 26)
	}
	return dkey, nil
}
