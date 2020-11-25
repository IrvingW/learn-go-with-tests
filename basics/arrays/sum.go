package arrays

// Sum return the sum of every item of an array
func Sum(numbers []int) (result int) {
	result = 0
	for _, number := range numbers {
		result += number
	}
	return
}

// SumAll return the sum of each slice item's sum
func SumAll(numbersToSum ...[]int) (result []int) {
	// check if input have same length
	length := 0
	for index, numbers := range numbersToSum {
		if index == 0 {
			length = len(numbers)
		} else if len(numbers) != length {
			println("InputException: input slices have different length")
			return
		}
	}
	for i := 0; i < length; i++ {
		sum := 0
		for _, numbers := range numbersToSum {
			sum += numbers[i]
		}
		result = append(result, sum)
	}
	return
}
