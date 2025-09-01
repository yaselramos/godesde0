package middlewares

import (
	"fmt"
)

func Sumar(a, b int) int {
	return a + b
}
func Restar(a, b int) int {
	return a - b
}

func Multiplicar(a, b int) int {
	return a * b
}
func Dividir(a, b int) int {
	if b == 0 {
		return 0 // Manejo simple de división por cero
	}
	return a / b
}

func MiMiddeleware() {
	fmt.Println("Este es un middleware de ejemplo")
	resultado := operacionesMidd(Sumar)(5, 3)
	fmt.Println("Resultado de la suma:", resultado)
	resultado = operacionesMidd(Restar)(10, 4)
	fmt.Println("Resultado de la resta:", resultado)
	resultado = operacionesMidd(Multiplicar)(6, 7)
	fmt.Println("Resultado de la multiplicación:", resultado)
	resultado = operacionesMidd(Dividir)(20, 4)
	fmt.Println("Resultado de la división:", resultado)
}
func operacionesMidd(f func(int, int) int) func(int, int) int {

	return func(i1, i2 int) int {
		fmt.Println("Iniciando operación...")
		return f(i1, i2)
	}
}
