package util

import (
	"encoding/base64"
	"log"
)

func Base64EncodeStr(raw string) string {
	encodeString := base64.StdEncoding.EncodeToString([]byte(raw))
	return encodeString
}

func Base64EncodeBytes(rawBytes []byte) string {
	encodeString := base64.StdEncoding.EncodeToString(rawBytes)
	return encodeString
}

func Base64DecodeStr(b64 string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(decodeBytes), nil
}
