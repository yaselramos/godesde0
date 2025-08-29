package teclado

import (
	"bufio"   // libriería para leer entradas
	"fmt"     // runtime es un paquete que maneja el sistema operativo
	"os"      // libriería systema operativo
	"strconv" // libriería para convertir tipos de datos
)

var numer01 int
var numer02 int
var leyenda string
var err error

func IngresoDatos() {
	fmt.Println("Ingrese el primer número")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numer01, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error en la conversión", err.Error())
		}
	}
	fmt.Println("El primer número es:", numer01)

	fmt.Println("Ingrese el 2 número")
	scanner1 := bufio.NewScanner(os.Stdin)
	if scanner1.Scan() {
		numer02, err = strconv.Atoi(scanner1.Text())
		if err != nil {
			fmt.Println("Error en la conversión", err.Error())
		}
	}
	fmt.Println("El primer número es:", numer02)

	fmt.Println("Ingrese la leyenda")
	scanner2 := bufio.NewScanner(os.Stdin)
	if scanner2.Scan() {
		leyenda, err = scanner2.Text(), nil
		if err != nil {
			fmt.Println("Error en la conversión", err.Error())
		}
	}
	fmt.Println(leyenda, numer01+numer02)
}
