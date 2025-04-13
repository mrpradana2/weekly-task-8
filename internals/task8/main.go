package task8

import "sync"

func PrintProduct(productName string, price string, channel chan map[string]string, wgx *sync.WaitGroup) {
	defer wgx.Done()
	newProduct := map[string]string{
		"Name Product": productName,
		"Price":        price,
	}
	channel <- newProduct
}