package variables

import "fmt"

// para llamar a una funcion de otro paquete, el nombre de la función debe comenzar con mayúscula
// si comienza con minúscula, es una función privada del paquete
func MuestraEnteras() {
	var intcomun int
	int32 := int32(3)
	int64 := int64(100)
	fmt.Println("intcomun:", intcomun)
	fmt.Println("int32:", int32)
	fmt.Println("int64:", int64)
}
