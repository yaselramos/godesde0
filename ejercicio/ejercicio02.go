package ejercicio

import (
	"bufio"   // libriería para leer entradas
	"fmt"     // runtime es un paquete que maneja el sistema operativo
	"os"      // libriería systema operativo
	"strconv" // libriería para convertir tipos de datos
)

var numer01 int
var err error
var resultado string

func TablaMul() string {

	for {
		fmt.Println("Ingrese el  número")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			numer01, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error vuelve a escribir el numero", err.Error())
				continue
			} else {
				break
			}

		}
	}

	for i := 0; i <= 10; i++ {
		resultado += fmt.Sprintf("%d x %d = %d\n", numer01, i, numer01*i)
	}
	return resultado + "\n"
}
