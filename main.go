package main

import (
	"fmt"

	"github.com/ddessilvestri/godesde0/variables"
)

func main() {
	variables.ShowIntegers()
	variables.RestVariables()
	state, text := variables.ConvertToText(1588)
	fmt.Println("state", state)
	fmt.Println("text", text)
}
