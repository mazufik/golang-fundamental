package calculation

func Add(number int, numberTwo int) int {
	return add(number, numberTwo)
}

/*
  perbedaan penamaan fungsi antara huruf depan kafital sama yang kecil adalah, kalau kafital menandakan
  fungsi itu dapat di panggil global dan yang kecil cuman bisa dipanggil dalam file/package yang sama.
*/

func add(number int, numberTwo int) int {
	return number + numberTwo
}
