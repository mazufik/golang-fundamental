package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*
		Proses transfer data pada channel secara default dilakukan dengan cara un-buffered, atau tidak di-buffer di memori. Ketika terjadi
		proses kirim data via channel dari sebuah goroutine, maka harus ada goroutine lain yang menerima data dari channel yang sama, dengan
		proses serah-terima yang bersifat blocking. Maksudnya, baris kode setelah kode pengiriman dan penerimaan data tidak akan di proses
		sebelum proses serah-terima itu sendiri selesai.
		Buffered channel sedikit berbeda. Pada channel jenis ini, ditentukan angka jumlah buffer-nya. Angka tersebut menjadi penentu jumlah
		data yang bisa dikirimkan bersamaan. Selama jumlah data yang dikirim tidak melebihi jumlah buffer, maka pengiriman akan berjalan
		asynchronous (tidak blocking).
		Proses pengiriman data pada buffered channel adalah asynchronous ketika jumlah data yang dikirim tidak melebihi batas buffer. Namun
		pada bagian channel penerimaan data selalu bersifat synchronous.
	*/

	runtime.GOMAXPROCS(2)

	messages := make(chan int, 3)

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}
}
