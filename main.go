package main

import (
	"fmt" // runtime es un paquete que maneja el sistema operativo
	"godesde0/ejercicio"
	"runtime"
)

func main() {
	//estado, testo, text := variables.ConviertoTexto(4343)
	//fmt.Println(estado, testo, text)

	if os := runtime.GOOS; os == "Windows_NT" || os == "OS X" {
		fmt.Println("El sistema operativo es:", os)
	} else {
		fmt.Println("El sistema operativo no es Windows ni linux")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("El sistema operativo es Windows")
	case "darwin":
		fmt.Println("El sistema operativo es darwin")
	case "linux":
		fmt.Println("El sistema operativo es Linux")
	default:
		fmt.Printf(" %s \n", os)
	}

	txt, text1 := ejercicio.Valores("150")
	fmt.Println(txt, text1)

}
