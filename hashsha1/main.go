package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	/*
		Hash adalah algoritma enkripsi untuk mengubah text menjadi deretan karakter acak. Jumlah karakter hasil hash selalu sama. Hash
		termasuk one-way encryption, membuat hasil dari hash tidak bisa dikembalikan ke text asli.
		SHA1 atau Secure Hash Algorithm 1 merupakan salah satu algoritma hashing yang sering digunakan untuk enkripsi data. Hasil dari
		sha1 adalah data dengan lebar 20 byte atau 160 bit, biasa ditampilkan dalam bentuk bilangan heksadesimal 40 digit.
	*/

	var text = "this is secret"

	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println("Original: ", text)
	fmt.Println("Hashed: ", encryptedString)
}
