package task2

import "fmt"

func DurationFilm(input uint) string {
	durationFilm := []uint{1, 7, 3, 4, 8, 9}
	for _, v := range durationFilm {
		for _, x := range durationFilm {
			if v+x == input {
				return fmt.Sprintf("%v dan %v", v, x);
			} 
		}
	}
	return "Tidak ada durasi film yang cocok"
}