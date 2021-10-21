package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	sentence := "Is your life you can play with it, you make you bed you 'gon lay in it, do ya thang just be safe with it,triple bucks in the state prison, blue laces in my blue chucks, ain't never gave two fucks, BET I chunk the hood up nigga, Asking If That Nipp hood or what"

	fmt.Println("Lenght Of The Text", len(Encodeb64(sentence)))
	fmt.Println(Encodeb64(sentence))
	fmt.Println("---------------------------------------------")
	fmt.Println("Lenght Of The Text", len(Decodeb64(Encodeb64(sentence))))
	fmt.Println(Decodeb64(Encodeb64(sentence)))
}

func Encodeb64(s string) string {
	// Set The Character That Can Contain The Encoding
	encodeStandard := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	encode := base64.NewEncoding(encodeStandard).EncodeToString([]byte(s))

	return encode
}

func Decodeb64(s string) string {
	d, err := base64.StdEncoding.DecodeString(s)

	handleErr(err)

	decode := string(d)

	return decode
}
func handleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
