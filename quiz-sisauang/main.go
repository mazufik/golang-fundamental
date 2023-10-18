package main

import "fmt"

func main() {
	var money int = 20
	items := map[string]int{"apel": 2, "pisang": 4, "jeruk": 6}

	for key, item := range items {
		fmt.Println("------------------------------------------------------")
		fmt.Println("Anda memiliki ", money, " dolar di dompet Anda")
		fmt.Println("Harga setiap ", key, " ", item, " dolar")

		var input_count int
		fmt.Print("Mau berapa ", key, " ? ")
		fmt.Scanf("%d", &input_count)
		fmt.Println("Anda akan membeli ", input_count, " ", key)

		total_price := item * input_count
		fmt.Println("Harga total adalah ", total_price, " dolar")

		if money >= total_price {
			fmt.Println("Anda telah membeli ", input_count, " ", key)
			money -= total_price

			if money == 0 {
				fmt.Println("Dompet Anda kosong")
				break
			}
		} else {
			fmt.Println("Uang Anda tidak mencukupi")
			fmt.Println("Anda tidak dapat membeli ", key, " sebanyak itu")
			break
		}
	}

	fmt.Println("Uang Anda tinggal ", money, " dolar")

}
