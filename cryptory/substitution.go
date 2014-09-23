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
var monoMap = map[rune]rune{'A': 'Z','B': 'W','C': 'X','D': 'Y','E': 'O','F': 'P','G': 'R','H': 'Q','I': 'A','J': 'H','K': 'C','L': 'B','M': 'E','N': 'S','O': 'U','P': 'T','Q': 'V','R': 'F','S': 'G','T': 'J','U': 'L','V': 'K','W': 'M','X': 'N','Y': 'D','Z': 'I'}

// Character Map for MonoAlphabetic Decryption
var monoMapRev = map[rune]rune{'Z': 'A','W': 'B', 'X': 'C', 'Y': 'D', 'O': 'E', 'P': 'F', 'R': 'G', 'Q': 'H', 'A': 'I', 'H': 'J', 'C': 'K', 'B': 'L', 'E': 'M', 'S': 'N', 'U': 'O', 'T': 'P', 'V': 'Q', 'F': 'R', 'G': 'S', 'J': 'T', 'L': 'U', 'K': 'V', 'M': 'W', 'N': 'X', 'D': 'Y', 'I': 'Z'}
