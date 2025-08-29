package file

import (
	"fmt"
	"godesde0/ejercicio"
	"os"
)

var file_dir string = "./file/txt/tabla.txt"

func File() {
	resultado := ejercicio.TablaMul()
	archivo, err := os.Create(file_dir)
	if err != nil {
		fmt.Println("Error al crear el archivo", err.Error())
		return
	}
	fmt.Fprintln(archivo, resultado)
	archivo.Close()
}
func SumaTabla() {
	resultado := ejercicio.TablaMul()
	if !Append(file_dir, resultado) {
		fmt.Println("Error al concatenar el archivo")
	}

}

func Append(file_dir string, resultado string) bool {
	archivo, err := os.OpenFile(file_dir, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return false
	}
	_, err = archivo.WriteString(resultado)
	if err != nil {
		return false
	}
	defer archivo.Close()
	return true
}

func ReadFile() string {
	archivo, err := os.ReadFile(file_dir)
	if err != nil {
		fmt.Println("Error al leer el archivo", err.Error())
		return ""
	}
	return string(archivo)
}
