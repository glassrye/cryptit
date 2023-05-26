package main

import (
	"fmt"

	"github.com/glassrye/cryptit/decrypt"
	"github.com/glassrye/cryptit/encrypt"
)

func main() {
		enc := fmt.Sprint(encrypt.Nimbus("Mark"))
		fmt.Println(enc)
		fmt.Println(decrypt.Nimbus(enc))
}