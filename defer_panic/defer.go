package defer_panic

import (
	"fmt"
	"log"
)

func Dedefer() {
	fmt.Println("Iniciando la función")
	defer fmt.Println("Esta línea se ejecuta al final de la función")
	fmt.Println("Ejecutando otras operaciones")
}

func EjemploPanic() {
	fmt.Println("Iniciando la función")
	panic("Ocurrió un error crítico") // Genera un panic aborto de la ejecución
	fmt.Println("Esta línea no se ejecutará debido al panic")
}

func EjemploRecover() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Recuperado del panic:", r)
		}
	}()
	fmt.Println("Iniciando la función")
	panic("Ocurrió un error crítico")
	fmt.Println("Esta línea no se ejecutará debido al panic")
}
