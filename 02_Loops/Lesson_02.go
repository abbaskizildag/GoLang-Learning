package main

import "fmt"

func main() {
	//go'da sadece tek bir döngümüz var o da for. fakat 3 farklı senaryosu mevcut.

	//1.kullanımı döngüde kullanacak değişken, kaça kadar döngünün devam etmesi ve döngünün artışı şeklinde
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//aşağıdaki gibide kullanılabilir
	k := 0
	for ; k < 5; k++ {
		fmt.Println(k)
	}

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

	//3.kullanım sonsuz döngü- Infinite Loop diye geçiyor.
	/*for {
		println("sonsuz bir döngü")
	}
	*/

	//continue döngünün başına dönmemizi sağlayan bir ifade. ekranda 0,3,6,9 görünmeyecek.
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	//break durumunda döngüden çıkıyor.
	for i := 0; i < 10; i++ {
		if i == 0 {
			break
		}
		fmt.Println(i)
	}

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
