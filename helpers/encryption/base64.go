package encryption

import "encoding/base64"

func Encode(data string) string {
	encodedString := base64.StdEncoding.EncodeToString([]byte(data))
	return encodedString
}

func Decode(data string) string {
	decodedByte, _ := base64.StdEncoding.DecodeString(data)
	decodedString := string(decodedByte)
	return decodedString
}
