package webserver

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hola, mundo! Este es un servidor web en Go.")
		http.ServeFile(w, r, "./index.html")
	})

	fmt.Println("Servidor web escuchando en http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
