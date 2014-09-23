package cryptory

import (
	"testing"
)

func TestCeasarEnc(t *testing.T) {
	var tt = ceasarValidTT()
	for key, value := range tt {
		encrypted, err := CeasarEnc(key)
		if err != nil {
			t.Fatalf("Error: %v \n", err)
		}
		if value != encrypted {
			t.Fatalf("Expected %v and got %v", value, encrypted)
		}
	}
}

func TestCeasarDec(t *testing.T) {
	var tt = ceasarValidTT()
	for key, value := range tt {
		decrypted, err := CeasarDec(value)
		if err != nil {
			t.Fatalf("Error: %v \n", err)
		}
		if key != decrypted {
			t.Fatalf("Expected %v and got %v", key, decrypted)
		}
	}
}

func TestMonoAlphaEnc(t *testing.T) {
	var tt = monoValidTT()
	for key, value := range tt {
		encrypted, err := MonoAlphaEnc(key)
		if err != nil {
			t.Fatalf("Error: %v \n", err)
		}
		if value != encrypted {
			t.Fatalf("Expected %v and got %v", value, encrypted)
		}
	}
}

func TestMonoAlphaDec(t *testing.T) {
	var tt = monoValidTT()
	for key, value := range tt {
		decrypted, err := MonoAlphaDec(value)
		if err != nil {
			t.Fatalf("Error: %v \n", err)
		}
		if key != decrypted {
			t.Fatalf("Expected %v and got %v", key, decrypted)
		}
	}
}

func ceasarValidTT() map[string]string {
	return map[string]string{
		"HELLO": "KHOOR",
	}
}

func ceasarInvalidTT() map[string]string {
	return map[string]string{
		"W0rLd":    "Z3uOg",
		"3, 32, 5": "6/#65/#8",
	}
}

func monoValidTT() map[string]string {
	return map[string]string{
		"HELLO": "QOBBU",
	}
}

func monoInvalidTT() map[string]string {
	return map[string]string{
		"W0rLd":    "Z3uOg",
		"3, 32, 5": "6/#65/#8",
	}
}
