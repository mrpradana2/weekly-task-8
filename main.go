package main

import (
	"fmt"
	"sync"
	"weekly-8/internals/task1"
	"weekly-8/internals/task2"
	"weekly-8/internals/task3"
	"weekly-8/internals/task4"
	"weekly-8/internals/task5"
	"weekly-8/internals/task6"
	"weekly-8/internals/task7"
	"weekly-8/internals/task8"
)

func main() {

	// task 1
	task1.PrintSegitiga(4)

	// taks 2
	fmt.Println(task2.DurationFilm(7))

	// task 3
	fmt.Println(task3.Pembulatan(4.234))

	// task 4
	deret := task4.DeretBilangan{N: 40}
	fmt.Println("Deret bilangan:", deret.CetakDeretBilangan())
	fmt.Println("Deret bilangan ganjil:", deret.Ganjil())
	fmt.Println("Deret bilangan genap:", deret.Genap())
	fmt.Println("Deret bilangan prima:", deret.Prima())
	fmt.Println("Deret bilangan fibonacci:", deret.Fibonacci())
	
	// task 5
	kubus := &task5.Kubus{}
	SetSisi(kubus, 5)
	HitungLuas(kubus)
	HitungKeliling(kubus)
	HitungVolume(kubus)

	// task 6
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	go task6.Sum(numbers, ch)
	total := <- ch
	fmt.Println(total)
	
	// task 7
	var wg sync.WaitGroup
	chn := make(chan []int, 2)
	wg.Add(2)
	go task7.Ganjil(numbers, chn, &wg)
	go task7.Genap(numbers, chn, &wg)
	wg.Wait()

	resultsGenap := <- chn
	resultGanjil := <- chn
	fmt.Println(resultsGenap)
	fmt.Println(resultGanjil)

	// task 8
	channel := make(chan map[string]string)
    var wgx sync.WaitGroup
	var mtx sync.Mutex
    wgx.Add(2)
    go func ()  {
        wgx.Wait()
        close(channel)
    }()

    go task8.PrintProduct("Sabun", "20000", channel, &wgx)
    go task8.PrintProduct("Shampo", "10000", channel, &wgx)
    for data := range channel {
        mtx.Lock()
        fmt.Println(data)
        mtx.Unlock()
    }
}

type hitung2d interface {
	Luas() float64
	Keliling() float64
	SetSisi(s float64)
}

type hitung3d interface {
	Volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

func HitungKeliling(h hitung)  {
	fmt.Println("Keliling:", h.Keliling())
}

func HitungLuas(h hitung) {
	fmt.Println("Luas:", h.Luas())
}

func HitungVolume(h hitung) {
	fmt.Println("Volume:", h.Volume())
}

func SetSisi(h hitung, s float64) {
	h.SetSisi(s)
}

