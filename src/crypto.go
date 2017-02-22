package main

import (
	"io"
	"crypto/md5"
	"encoding/hex"
)

func md5string(subject string) string {
	h := md5.New()
    io.WriteString(h, subject)
    return hex.EncodeToString(h.Sum(nil))
}