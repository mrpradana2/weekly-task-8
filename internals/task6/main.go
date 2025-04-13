package task6

func Sum(numbers []int, result chan int) {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	result <- sum
}