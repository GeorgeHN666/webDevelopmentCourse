package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	test1 := hMAC("george1029@gmail.com")
	fmt.Println(test1)
	test2 := hMAC("george1o29@gmail.com")
	fmt.Println(test2)
}

func hMAC(text string) string {
	hash := hmac.New(sha256.New, []byte("Key"))

	io.WriteString(hash, text)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// This Tool Called HMAC that stands for Hash_Message_Autentication_Code Will Alllow us create a hash specific for a piece of information that we want to encode That We Can Compare For An autentication Method
