package main

import "fmt"

//sum isminde bir func tanımladık. bu arkadaş x ve y tipinde değer alır. geriye x+y'yi toplarak int döner dedik.
func sum(x, y int) int {
	return x + y
}

//string değer alan ve geriye bir şey dönmeyen bir func yazdık
func hello(x string) {
	fmt.Println("func içerisinde geliyorum", x)
}

//hem string hem int
func numberstring(x int, y string) {
	fmt.Println(x, y)
}

//birden fazla değer dönebiliyoruz. string, int fark etmiyor.
func tworeturn(x int, y string) (int, string) {
	return x + 5, y + "deneme"
}

func main() {

	fmt.Println(sum(5, 10)) //ekrana yazdırdık
	hello("merhaba")
	numberstring(5, "test")

	i, j := tworeturn(5, "main'den gelen")
	println(i, j)

}
