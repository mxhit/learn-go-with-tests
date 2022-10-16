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
	// getting number of slices passed in the argument
	noOfSlices := len(slices)
	// creating a slice with length = noOfSlices
	sums := make([]int, noOfSlices)

	// iterating over all the slices passed in the argument
	for i, slice := range slices {
		// adding the value returned by Sum in the respective index
		sums[i] = Sum(slice)
	}

	return sums
}
