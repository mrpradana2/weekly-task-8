package task1

import "fmt"

func PrintSegitiga(maxRow int) {
	for i := 1; i <= maxRow; i++ {
		var star string
		for j := maxRow * 2; j >= i; j-- {
			if (maxRow * 2) - i < j {
				star += " "
			} else {
				star += "*"
			}
		}
		fmt.Println(star)
	}
}