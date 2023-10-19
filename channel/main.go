package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*
		Channel digunakan untuk menghubungkan goroutine satu dengan goroutine lain. Mekanisme yang dilakukan adalah serah-terima data
		lewat channel tersebut. Dalam komunikasinya, sebuah channel difungsikan sebagai pengirim di sebuah goroutine, dan juga sebagai
		penerima di goroutine lainnya. Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous.
	*/

	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("Hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}
