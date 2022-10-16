package iteration

func Repeat(character string, repeatCount int) string {
	var repeated string

	if repeatCount <= 0 {
		return ""
	}

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
