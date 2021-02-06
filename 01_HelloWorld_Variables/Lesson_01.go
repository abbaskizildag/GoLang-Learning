package main

import "fmt"

func main() {
	fmt.Println("Hello World") //Bir programlama diline daha merhaba,

	//1.Değişken Tanımlama
	var firstname string
	firstname = "Ali"

	//Yukarıdaki değişkeni tek satırda tanımlama
	var lastname string = "Kizildag"

	//Birden fazla değişken tanımlama
	var name, surname string = "Yasin", "Demir"

	var age int = 29              //sayısal veri tanımlama int yerine rune tipide kullanılabilir
	var isTrue = true             //bool değer atama
	var floatNumber float32 = 8.5 //virgüllü değer atama

	//tip belirtmeden değişken atama
	city := "kars"

	fmt.Println(firstname, lastname)
	fmt.Println(name, surname)
	fmt.Println("Age=", age, "isTrue=", isTrue)
	fmt.Println(floatNumber)
	fmt.Println(city)

	//eğer tiplerini öğrenmek istersek int mi string mi bool mu
	fmt.Printf("firstname'ın değişken tipi %T içindeki %s değer\n age'in değişken tiği %T değeri %d", firstname, firstname, age, age)
}
