package funciones

import "fmt"

func Carculos() {

	var numero3 int = 10
	var numero4 int = 20
	suma := func(numer01 int, numer02 int) int {
		return numer01 + numer02 + numero3 + numero4

	}
	fmt.Println(suma(7, 8))

	suma = func(numer01 int, numer02 int) int {
		return numer01 + numer02*numero3

	}
	fmt.Println(suma(7, 8))
}
