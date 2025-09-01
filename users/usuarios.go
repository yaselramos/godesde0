package users

import (
	"fmt"
	"godesde0/modelos"
	"time"
)

func AltaUsuario() {
	user := new(modelos.User)
	//var user modelos.User
	user.AddUser(1, "Juan Perez", "yramosm@gmail.com", time.Now(), true)
	fmt.Println(user)
}
