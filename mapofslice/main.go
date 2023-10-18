package main

import "fmt"

func main() {
	students := []map[string]string{
		{
			"name":  "Agung",
			"score": "A",
		},
		{
			"name":  "Taufik",
			"score": "A",
		},
		{
			"name":  "Zaydan",
			"score": "B",
		},
		{
			"name":  "Mario",
			"score": "C",
		},
	}

	for _, value := range students {
		// fmt.Println("key : ", key, " value : ", value)
		fmt.Println("=============================")
		fmt.Println(value["name"], " - ", value["score"])
	}
}
