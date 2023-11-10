package main

import (
	"crypto/md5"
	"fmt"
)

func md5Encode(input string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(input))
	md5 := hash.Sum(nil)
	return fmt.Sprintf("%x", md5)
}
