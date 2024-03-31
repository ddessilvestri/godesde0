package main

import (
	"fmt"
	"runtime"

	"github.com/ddessilvestri/godesde0/ejercicios"
	"github.com/ddessilvestri/godesde0/variables"
)

func main() {
	num, texto := ejercicios.ConvNumerico("x")

	fmt.Println(num)
	fmt.Println(texto)

	estado, texto := variables.ConvertoaText(1858)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows")
	} else {
		fmt.Println("Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("Darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
