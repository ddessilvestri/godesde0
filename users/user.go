package users

import (
	"fmt"
	"time"

	"github.com/ddessilvestri/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u.Name)
	fmt.Println(u)
}
