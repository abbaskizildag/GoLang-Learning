package main

import "fmt"

func main() {
	fmt.Println("Hello World") //Bir programlama diline daha merhaba,

	//1.Değişken Tanımlama- Önce tanımlama yapıyoruz sonra değeri atıyoruz
	var firstname string
	firstname = "Ali"

	//Yukarıdaki değişkeni tek satırda tanımlama ve değer atama yapıyoruz
	var lastname string = "Kizildag"

	//Birden fazla değişken tanımlama ve değer ataması yapıyoruz.
	var name, surname string = "Yasin", "Demir"

	//2.değişken tanımlama sadece değişken adı ve değerini veriyoruz tipi kendi alıyor
	var age = 29          //sayısal veri tanımlama int yerine rune tipide kullanılabilir
	var isTrue = true     //bool değer atama
	var floatNumber = 8.5 //virgüllü değer atama

	//var age, isTrue, floatNumber = 29, true, 8.5 //tek satırda da tanımlayabiliriz.

	//3.değişken tanımlama. Bu en çok kullanacağımız. short declaration deniliyor.
	city := "kars"
	//tek satırda birden fazla yukarıdaki gibi 3.tanımlama yöntemiyle tanımlama yapabiliriz.
	//name,surname :="ali","veli"

	//Ekrana yazdırmak için Print, Println ve Printf kullanılıyor. Println direk alt satıra atıyor. Print aynı satırda yazdırıyor.
	//Printf ise fonksiyonun tipi, string ifade içerisinde nasıl görüntülenmesini istiyorsak onun için kullanıyoruz.
	fmt.Println(firstname, lastname)
	fmt.Println(name, surname)
	fmt.Println("Age=", age, "isTrue=", isTrue)
	fmt.Println(floatNumber)
	fmt.Println(city)

	//eğer tiplerini öğrenmek istersek int mi string mi. %T kullanıyoruz.
	fmt.Printf(" firstname'ın değişken tipi= %T içindeki değer = %s \n age'in değişken tipi= %T değeri= %d\n floatNumber'ın tipi= %T ve değeri= %f", firstname, firstname, age, age, floatNumber, floatNumber)
}
