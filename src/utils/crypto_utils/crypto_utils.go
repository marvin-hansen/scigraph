// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package crypto_utils

import (
	"encoding/base64"
)

func EncodeBase64(stringToEncode string) (encodedString string) {
	// https://gobyexample.com/base64-encoding
	encodedString = base64.StdEncoding.EncodeToString([]byte(stringToEncode))
	return encodedString
}
func DecodeBase64(stringToDecode string) (decodedString string) {
	bytes, _ := base64.StdEncoding.DecodeString(stringToDecode)
	decodedString = string(bytes)
	return decodedString
}
