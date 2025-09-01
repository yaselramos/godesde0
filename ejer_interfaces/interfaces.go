package ejer_interfaces

import (
	"fmt"
	"godesde0/interfaces"
)

func HumanosRespirar(humano interfaces.Humano) {
	humano.Respirar()
	fmt.Println("Soy un/a %s y estoy respirando \n", humano.Sexo())
}
