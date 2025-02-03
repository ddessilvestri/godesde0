package main

import (
	"fmt"
	"runtime"

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
}
