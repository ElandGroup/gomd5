package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(Sample("1234"))
	fmt.Println(GetMD5Hash("1234"))

}

func Sample(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))

}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return fmt.Sprintf("\n%X\n", hasher.Sum(nil)) //this is my choice
}
