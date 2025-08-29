package iteraciones

import (
	"fmt"
)

func IteracionesFor() {

	i := 0
	for i < 10 {
		fmt.Println("El valor de i es:", i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("El valor de i es:", i)
	}
}
