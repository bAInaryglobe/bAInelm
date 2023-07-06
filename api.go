package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	apiKey := hex.EncodeToString(key)
	fmt.Println(apiKey)
}
