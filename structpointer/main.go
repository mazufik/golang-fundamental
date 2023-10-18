package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

// Method
func (student *Student) graduate() {
	student.Name = student.Name + ", S.T"
}

func main() {
	// // pointer struct sebagai parameter
	// student := Student{1, "Zaydan Rahman", 3.87}
	//
	// fmt.Println(student.Name)
	//
	// graduate(&student)
	//
	// fmt.Println(student.Name)

	// Method dengan pointer receiver
	student := Student{1, "Zaydan Rahman", 3.87}
	fmt.Println(student.Name)

	student.graduate()

	fmt.Println(student.Name)
}

// // pointer struct sebagai parameter
// func graduate(student *Student) {
// 	student.Name = student.Name + ", S.T"
// }
