package task4

type DeretBilangan struct {
	N uint
}

func (d DeretBilangan) CetakDeretBilangan() []int {
	var deret []int
	for i := 1; i <= int(d.N); i++ {
		deret = append(deret, i)
	}
	return deret
}

func (d DeretBilangan) Ganjil() []int {
	var ganjil []int
	deretBilangan := d.CetakDeretBilangan()
	for _, angka := range deretBilangan {
		if angka%2 == 1 {
			ganjil = append(ganjil, angka)
		}
	}
	return ganjil
}

func (d DeretBilangan) Genap() []int {
	var genap []int
	deretBilangan := d.CetakDeretBilangan()
	for _, angka := range deretBilangan {
		if angka%2 == 0 {
			genap = append(genap, angka)
		}
	}
	return genap
}

func (d DeretBilangan) Prima() []int {
	var prima []int
	deretBilangan := d.CetakDeretBilangan()
	for _, angka := range deretBilangan {
		if angka < 2 {
			continue
		}

		check := true
		for i := 2; i*i <= angka; i++ {
			if angka%i == 0 {
				check = false
				break
			}
		}

		if check {
			prima = append(prima, angka)
		}
	}
	return prima
}

func (d DeretBilangan) Fibonacci() []int {
	n := d.N
	fibonacci := []int{0, 1}
	for range int(n) {
		length := len(fibonacci)
		newElement := fibonacci[length-2] + fibonacci[length-1]
		if newElement > int(n) {
			break
		}
		fibonacci = append(fibonacci, newElement)
	}
	return fibonacci
}