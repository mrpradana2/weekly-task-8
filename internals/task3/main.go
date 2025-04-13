package task3

import (
	"fmt"
	"math"
)

func Pembulatan(n float64) string {
	rounded := math.Round(n * 10) / 10
	return fmt.Sprintf("%0.2f", rounded)
}