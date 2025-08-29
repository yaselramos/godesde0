package main

import (
	"fmt"
	"godesde0/variables"
)

func main() {
	estado, testo, text := variables.ConviertoTexto(4343)
	fmt.Println(estado, testo, text)
}
