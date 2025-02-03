package main

import (
	"fmt"
	"runtime"

	"github.com/ddessilvestri/godesde0/excercises"
	"github.com/ddessilvestri/godesde0/iterations"
	"github.com/ddessilvestri/godesde0/variables"
)

func main() {
	variables.ShowIntegers()
	variables.RestVariables()
	state, text := variables.ConvertToText(1588)
	fmt.Println("state", state)
	fmt.Println("text", text)

	// os:= runtime.GOOS

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Linux or OS", os)
	} else {
		fmt.Println("should be windows: ", os)

	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("linux")
	case "darwin":
		fmt.Println("darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	exInt, exString, err := excercises.Excercise1("1500000")
	fmt.Println("exInt", exInt)
	fmt.Println("exInt", exString)
	fmt.Println("err", err)

	// keyboards.InputNumbers()

	for {
		fmt.Println("about to break")
		break
	}

	for i := 0; i < 2; i++ {
		fmt.Println("about to break in:", i)

	}
	iterations.Iterar()
	excercises.Excercise2()
}
