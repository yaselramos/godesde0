package arreglos_slice

import "fmt"

var slice []int                              //slice de enteros tamaño dinámico
var arreglo [10]int = [10]int{1, 2, 3, 4, 5} //arreglo de enteros tamaño fijo
func Slice() {
	slice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(slice)
	porcion := arreglo[5:] //rebanada del arreglo desde la posición 5 hasta el final
	fmt.Println(porcion)
	porcion2 := arreglo[1:4] //rebanada del arreglo desde la posición 1 hasta la 4 (sin incluir la 4)
	fmt.Println(porcion2)
	slice = append(slice, 17) //agrega un elemento al final del slice
	fmt.Println(slice)
	slice = append(slice, 19, 23, 29) //agrega varios elementos al final del slice
	fmt.Println(slice)
	slice2 := make([]int, 10) //crea un slice de enteros con capacidad para 10 elementos
	fmt.Println(slice2)
	copy(slice2, slice) //copia los elementos del slice en el slice2
	fmt.Println(slice2)

}

func Capacidad() {
	elementos := make([]int, 5, 20) //slice de enteros con capacidad para 5 elementos y 20 de tamaño máximo
	fmt.Printf("len: %d, Cap %d", len(elementos), cap(elementos))

	num := make([]int, 0, 0) //slice de enteros con capacidad para 0 elementos y 0 de tamaño máximo
	for i := 0; i < 100; i++ {
		num = append(num, i)
	}
	fmt.Printf("len: %d, Cap %d", len(num), cap(num))
}
