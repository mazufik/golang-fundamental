package main

import "fmt"

// interface
type BangunDatar interface {
	HitungLuas() int
}

// struct
type Persegi struct {
	Sisi int
}

// method
func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

// struct persegi panjang
type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

// method persegi Panjang
func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	/*
				Interface secara konsep dapat kita padankan dengan kontrak ada aturan mainnya (kontraknya)
				yang harus dipenuhi. di dalam interface berisikan kontrak berupa method. cara untuk memenuhi sebuah kontrak
		    adalah dengan membuat struct.
	*/
	persegiPanjang := PersegiPanjang{Panjang: 6, Lebar: 5}
	luas := SeberapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang : ", luas)

	persegi := Persegi{Sisi: 5}
	luas = SeberapaLuas(persegi)
	fmt.Println("Luas Persegi : ", luas)

}

// function
func SeberapaLuas(bangunData BangunDatar) int {
	return bangunData.HitungLuas()
}
