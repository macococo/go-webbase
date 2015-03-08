package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

func HandleError(err error) error {
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetGUID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
