package arraysslices

func Sum(numbers []int) int {
	sum := 0

	// range returns two values: the index and the
	// _ is a blank identifier since we do not want to use the index provided by range
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	var sums []int

	// iterating over all the slices passed in the argument
	for _, slice := range slices {
		// appending the value returned by Sum in sums slice
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int

	// iterating over all the slices passed in the argument
	for _, slice := range slices {
		if len(slice) > 1 {
			// creating a new slice by excluding the head i.e. 0th index
			tail := slice[1:]
			// appending the sum of tail
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}

	}

	return sums
}
