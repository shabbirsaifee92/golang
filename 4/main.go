package iteration

func Repeat(character string) string {
	var repeatedString string
	for i := 0; i < 5; i++ {
		repeatedString += character
	}
	return repeatedString
}
