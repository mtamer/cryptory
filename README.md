Cryptory
========

Implementation of Different Algos

## Substitution Algorithms

Ceasar's Cipher

``` go
func main() {
	fmt.Println(cryptory.CeasarEnc("HELLO"))
	// KHOOR
	fmt.Println(cryptory.CeasarDec("KHOOR"))
	// HELLO
}
```

MonoAlphabetic Cipher

``` go
func main() {
	fmt.Println(cryptory.MonoAlphaEnc("HELLO"))
	// QOBBU
	fmt.Println(cryptory.MonoAlphaDec("QOBBU"))
	// HELLO
}
```

Hill Cipher

``` go
func main() {
	orig := "HATS"
	fmt.Printf("Original Message: %v \n", orig)
	// HATS

	key, encrypted, err := cryptory.HillEnc(orig)
	if err != nil {
		fmt.Errorf("Error:", err)
	}
	fmt.Printf("Encrypted Message: %v \n", encrypted)
	// VOHY

	msg, err := cryptory.HillDec(encrypted, key)
	if err != nil {
		fmt.Errorf("Error:", err)
	}
	fmt.Printf("Decrypted Message: %v \n", msg)
	// HATS
}

# TODO

Hill Cipher - Remove hard coded encryption key
Hill Cipher - Ugly byte to int to byte conversions
Playfair Cipher