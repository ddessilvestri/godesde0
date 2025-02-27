package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.79
	Fecha = time.Now()

	fmt.Println("Nombre", Nombre)
	fmt.Println("Estado", Estado)
	fmt.Println("Sueldo", Sueldo)
	fmt.Println("Fecha ", Fecha)
}

func ConvertToText(number int) (bool, string) {
	// text := strconv.Itoa(number)
	return true, strconv.Itoa(number)
}
