package encrypt

// Nimbus takes a string and returns an encrypted version of it
func Nimbus(str string) string {
	encryptedStr := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode + 3)
		encryptedStr += character
	}


	return encryptedStr
}