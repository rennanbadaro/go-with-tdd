package iteration

const defaultRepetition = 5

func Repeat(character string, repetition int) string {
	if repetition <= 0 {
		repetition = defaultRepetition
	}

	var repeated string

	for i := 0; i < repetition; i++ {
		repeated += character
	}

	return repeated
}
