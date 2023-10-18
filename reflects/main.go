package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		Reflection adalah teknik untuk inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkan memanipulasinya. Cakupan
		informasi yang bisa didapatkan lewat reflection sangat luas, seperti melihat struktur variabel, tipe, nilai pointer, dan banyak lagi.
		Go menyediakan package reflect , berisikan banyak sekali fungsi untuk keperluan reflection. Pada chapter ini, kita akan belajar
		tentang dasar penggunaan package tersebut.
	*/

	number := 23
	reflectValue := reflect.ValueOf(number)

	fmt.Println("tipe variabel : ", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel : ", reflectValue.Int())
	}
}
