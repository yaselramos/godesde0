package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)    //crea un mapa con clave string y valor string
	paises["Mexico"] = "CDMX"            //agrega un elemento al mapa
	paises["Colombia"] = "Bogota"        //agrega un elemento al mapa
	paises["Argentina"] = "Buenos Aires" //agrega un elemento al mapa
	fmt.Println(paises)                  //imprime el mapa
	fmt.Println(paises["Colombia"])      //imprime el valor del elemento con clave "Colombia"

	ciudades := map[string]int{ //crea un mapa con clave string y valor string
		"Mexico":    36,
		"Colombia":  78,
		"Argentina": 45,
	}
	fmt.Println(ciudades)
	for clave, valor := range ciudades { //recorre el mapa
		fmt.Printf("Clave: %s, Valor: %d \n", clave, valor)
	}

	delete(ciudades, "Argentina") //elimina el elemento con clave "Argentina"

	fmt.Println(ciudades)

	valor, existe := ciudades["Colombia"] //verifica si el elemento con clave "Colombia" existe
	if existe {
		fmt.Println("El valor es:", valor, existe)
	} else {
		fmt.Println("El elemento no existe")
	}
}
