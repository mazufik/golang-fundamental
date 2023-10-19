package main

import (
	"fmt"
	"io"
	"os"
)

var path = "/home/zaydan/PORTFOLIO/DEVELOPMENT/GITHUB/GOLANG-PROJECT/GOLANG-DASAR/file/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

// membuat file
func createFile() {
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("===> file berhasil dibuat", path)
}

// menulis data ke file
func writeFile() {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString("hello\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) {
		return
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("===> file berhasil diisi")
}

// membaca file
func readFile() {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}

	fmt.Println("===> file berhasil dibaca")
	fmt.Println(string(text))
}

// menghapus file
func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("===> file berhasil dihapus")
}

func main() {
	var options = [4]string{"1. Buat File", "2. Mengedit File", "3. Membaca File", "4. Hapus File"}
	for _, option := range options {
		fmt.Println(option)
	}

	var option int
	fmt.Print("Type for option: ")
	fmt.Scanln(&option)

	switch option {
	case 1:
		createFile()
	case 2:
		writeFile()
	case 3:
		readFile()
	case 4:
		deleteFile()
	}
}
