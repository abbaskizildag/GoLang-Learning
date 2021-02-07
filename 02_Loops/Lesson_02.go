package main

import "fmt"

func main() {
	//go'da sadece tek bir döngümüz var o da for. fakat 3 farklı senaryosu mevcut.
	var x, y, z int = 0, 1, 0

	//1.örnek for kullanımı - fibonacci sayı sistemini for ile bulduk. her sayı öncekinin toplamından oluşur. 8 basamağa kadar yaptık.
	for i := 0; i < 8; i++ {
		z = y
		y = x + y
		x = z
		fmt.Printf("%d\t", z) //t tab yani boşluk bırakıyor.
	}

	//2.örnek kullanımı(bildiğiniz while döngüsü şeklinde)
	var a, b int = 15, 0
	for a > b {
		b++
		fmt.Printf("%d\n ", b)
	}

	//3.kullanım sonsuz döngü
	/*for {
		println("sonsuz bir döngü")
	}
	*/
	//genel örnek. iç içe for. asal sayıları bulma
	var i, j int

	for i = 2; i < 20; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d asal sayı\n", i)
		}
	}
}
