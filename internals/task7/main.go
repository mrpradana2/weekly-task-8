package task7

import "sync"

func Ganjil(numbers []int, result chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var filter []int
	for _, number := range numbers {
		if number%2 == 1 {
			filter = append(filter, number)
		}
	}
	result <- filter
}

func Genap(numbers []int, result chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var filter []int
	for _, number := range numbers {
		if number%2 == 0 {
			filter = append(filter, number)
		}
	}
	result <- filter
}