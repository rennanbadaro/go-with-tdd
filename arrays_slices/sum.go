package arraySlices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(collections ...[]int) []int {
	var sums []int

	for _, collection := range collections {
		sums = append(sums, Sum(collection))
	}

	return sums
}

func SumAllTails(collections ...[]int) []int {
	var tailSum []int

	for _, collection := range collections {
		if len(collection) > 1 {
			tail := collection[1:]

			tailSum = append(tailSum, Sum(tail))
		} else {
			tailSum = append(tailSum, 0)
		}
	}

	return tailSum
}
