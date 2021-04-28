package convert

import "encoding/base64"

func EncodeBase64(byts []byte) string {
	return base64.StdEncoding.EncodeToString(byts)
}
func DecodeBase64(str string) []byte {
	byts, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return []byte{}
	}
	return byts
}
