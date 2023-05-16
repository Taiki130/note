package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func sha1Checksum(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	checksum := sha1Checksum("Hello World!")
	fmt.Println(checksum)
}
