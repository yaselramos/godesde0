package main

import (
	"godesde0/gorutines"
)

func main() {
	//estado, testo, text := variables.ConviertoTexto(4343)
	//fmt.Println(estado, testo, text)

	/*if os := runtime.GOOS; os == "Windows_NT" || os == "OS X" {
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

	txt, text1 := ejercicio.Valores("7")
	fmt.Println(txt, text1)

	teclado.IngresoDatos()
	iteraciones.IteracionesFor()
	resultado := ejercicio.TablaMul()
	fmt.Println(resultado)
	file.File()
	file.SumaTabla()
	fmt.Println(file.ReadFile())
	funciones.Carculos()
	funciones.LlamarClosures()
	funciones.Exponencial(2)
	arreglos_slice.Arreglos()
	mapas.MostrarMapas()
	users.AltaUsuario()
	Persona1 := new(modelos.Hombre)
	ejer_interfaces.HumanosRespirar(Persona1)
	Persona2 := new(modelos.Mujer)
	ejer_interfaces.HumanosRespirar(Persona2)*/
	//defer_panic.Dedefer()
	//defer_panic.EjemploPanic()
	//defer_panic.EjemploRecover()
	/*go gorutines.MiNombreLentoo("Yasel Ramos Medina") //Con la palabra go se crea una gorutina

	fmt.Println("Hola")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Fin del programa")*/

	canal := make(chan bool) // Crear un canal para sincronización
	go gorutines.MiNombreLento("Yasel Ramos Medina", canal)
	estado := <-canal // Esperar a que la gorutina termine
	if estado {
		println("\nLa gorutina ha terminado")
	}
}
