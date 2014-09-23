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

# TODO

Playfair Cipher