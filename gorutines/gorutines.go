package gorutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLentoo(nombre string) {
	letra := strings.Split(nombre, "")
	for _, letra := range letra {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print(letra)
	}
}
func MiNombreLento(nombre string, canal chan bool) {
	letra := strings.Split(nombre, "")
	for _, letra := range letra {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print(letra)
	}
	canal <- true // Indica que la gorutina ha terminado
}
