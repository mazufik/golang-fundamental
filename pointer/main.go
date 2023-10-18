package main

import "fmt"

func main() {
	/*
		pointer adalah sebuah variabel yang isinya adalah suatu alamat dari suatu nilai,
		jadi yang disimpan itu bukan nilainya tapi alamat memory di komputer dimana nilai itu
		disimpan
	*/

	// numberA := 3
	// numberB := &numberA // proses reverensing, yaitu proses mengambil alamat memory nilai numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)

	/*
	   proses direverynsing adalah proses mengembalikan dari alamar memory ke nilainya
	*/
	// fmt.Println(*numberB)

	/*
		numberA dan numberB memiliki keterkaitan satu saya lainnya, yaitu keterkaitan alamat penyimpanan di memory,
		apabila nilai deliverynsing numberB dirubah maka nilai numberA akan ikut terubah
	*/

	// *numberB = 10

	// fmt.Println(*numberB)
	// fmt.Println(numberA)

	// // cara lainnya
	// var numberA int = 5
	// var numberB *int = &numberA
	//
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)
	//
	// numberA = 20
	//
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// Studi Kasus Penggunaan pointer
	number := 5
	fmt.Println("Alamat memory : ", &number)
	fmt.Println("Nilai awal : ", number)

	change(&number, 100)

	fmt.Println("Nilai akhir : ", number)
	fmt.Println("Alamat memory : ", &number)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("alamat memory : ", old)
	fmt.Println("di dalam function : ", *old)
}
