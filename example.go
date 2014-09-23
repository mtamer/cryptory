package main

import (
	"fmt"
	"github.com/collinglass/cryptory/cryptory"
)

func main() {
	fmt.Println(cryptory.CeasarEnc("hello"))
	fmt.Println(cryptory.CeasarEnc("W0rLd"))
	fmt.Println(cryptory.CeasarEnc("3, 32, 5"))

	fmt.Println(cryptory.MonoAlphaEnc("hello"))
	fmt.Println(cryptory.MonoAlphaEnc("W0rLd"))
	fmt.Println(cryptory.MonoAlphaEnc("3, 32, 5"))

}
