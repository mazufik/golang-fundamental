package main

import "fmt"

func main() {
	// soal 1
	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	//
	// var total int
	// length := len(scores)
	//
	// for _, value := range scores {
	// 	total += value
	// }
	//
	// average := float64(total) / float64(length)
	//
	// fmt.Println(average)

	// soal 2
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var goodScores []int

	for _, value := range scores {
		if value >= 90 {
			goodScores = append(goodScores, value)
		}
	}

	fmt.Println(goodScores)

}
