package main

import "fmt"

func main() {
	score := 79
	var grade string

	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score < 80 {
		grade = "C"
	} else if score <= 60 {
		grade = "D"
	} else if score <= 50 {
		grade = "E"
	} else {
		grade = "NULL"
	}

	fmt.Println(grade)
}
