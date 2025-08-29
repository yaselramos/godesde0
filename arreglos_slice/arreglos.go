package arreglos_slice

import "fmt"

var tabla [10]int    //arreglo de 10 enteros vectores tamaño fijo
var matriz [4][4]int //matriz de 4x4
func Arreglos() {
	for i := 0; i < len(tabla); i++ {
		tabla[i] = i * 2
	}
	fmt.Println(tabla)
	matriz[0][2] = 1
	fmt.Println(matriz)

}
