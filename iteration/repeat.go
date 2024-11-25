package iteration

func Repeat(char string, repeats int) string {
	var repeated string
	for i := 0; i < repeats; i++ {
		repeated += char
	}
	return repeated
}
