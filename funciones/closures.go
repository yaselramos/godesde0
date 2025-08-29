package funciones

import "fmt"

func tabla(varlor int) func() int {
	numero := varlor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
func LlamarClosures() {
	tablaDel := 2
	miTabla := tabla(tablaDel)
	for i := 1; i < 10; i++ {
		fmt.Println(miTabla())
	}
}
