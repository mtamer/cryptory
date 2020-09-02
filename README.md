你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
```

# TODO

Hill Cipher - Remove hard coded encryption key
Hill Cipher - Ugly byte to int to byte conversions
Playfair Cipher
