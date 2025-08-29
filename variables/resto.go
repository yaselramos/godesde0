package variables

import (
	"fmt"     // fmt es un paquete que permite imprimir en consola
	"strconv" // strconv es un paquete que convierte tipos de datos
	"time"    // time es un paquete que maneja fechas y horas
)

var Nombre string
var Estado bool
var Ent int
var Sueldo float32
var fecha time.Time

func RestoVariables() {
	Nombre = "Golang"
	Estado = true
	Ent = 100
	Sueldo = 3.1416
	fecha = time.Now().Local()
	fmt.Println("Nombre:", Nombre)
	fmt.Println("Estado:", Estado)
	fmt.Println("Entero:", Ent)
	fmt.Println("Flotante:", Sueldo)
	fmt.Println("Fecha:", fecha)
}

func ConviertoTexto(numero int) (bool, string, string) {
	text := strconv.Itoa(numero) // convierte un entero a texto
	texto := fmt.Sprintf("El número es: %d", numero)
	return true, texto, text
}
