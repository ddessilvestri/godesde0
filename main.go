package main

import (
	"fmt"

	"github.com/ddessilvestri/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvertoaText(1858)
	fmt.Println(estado)
	fmt.Println(texto)
}
