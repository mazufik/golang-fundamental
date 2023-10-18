package main

import "fmt"

func main() {
	// cara pertama
	// var myMap map[string]int
	// myMap = map[string]int{}
	//
	// myMap["Ruby"] = 9
	// myMap["Javascript"] = 8
	// myMap["Go"] = 10
	//
	// fmt.Println(myMap)
	// fmt.Println(myMap["Go"])

	// cara kedua
	// myMap := map[string]string{"Ruby": "is beautiful", "Go": "is fast"}
	//
	// fmt.Println(myMap)
	// fmt.Println(myMap["Go"])

	// meanmpilkan map dengan looping
	myMap := map[string]string{
		"Ruby":       "is beautiful",
		"Go":         "is fast",
		"Javascript": "is hype",
		"Python":     "is stable",
	}

	for key, value := range myMap {
		fmt.Println("Key : ", key, " Value : ", value)
	}

	// menghapus isi map
	fmt.Println("================")

	delete(myMap, "Python")

	for key, value := range myMap {
		fmt.Println("Key : ", key, " Value : ", value)
	}
}
