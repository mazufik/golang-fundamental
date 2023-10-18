package main

import "fmt"

func main() {
	var scores = []int{100, 85, 75, 60, 90, 93, 80, 91}
	var avg = calculate(scores...)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
}

func calculate(scores ...int) float64 {
	var total int = 0
	for _, score := range scores {
		total += score
	}

	var avg = float64(total) / float64(len(scores))
	return avg
}
