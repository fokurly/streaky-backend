package utils

import (
	"crypto/md5"
	"fmt"
)

func HashPassword(pass string) string {
	h := md5.New()
	h.Write([]byte(pass))
	return fmt.Sprintf("%x", h.Sum(nil))
}
