package ejercicio

import (
	"strconv"
)

func Valores(text string) (int, string) {
	text1, err := strconv.Atoi(text) // convierte un texto a entero
	if err != nil {
		// Si hay error, puedes retornar 0 o cualquier valor por defecto
		return 0, text
	}
	if text1 < 100 {
		return text1, "El número es menor a 100"
	} else {
		return text1, "El número es mayor o igual a 100"
	}

}
