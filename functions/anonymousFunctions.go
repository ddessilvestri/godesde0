package functions

import "fmt"

func Calculates() {
	var numero3 int = 32
	var numero4 int = 46

	suma := func(num1 int, num2 int) int {
		return num1 + num2 + numero3 + numero4
	}
	fmt.Println(suma(1, 2))

}
