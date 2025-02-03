package excercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error

func Excercise2() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingresar un numero")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
	}
	fmt.Println("numero", numero1)
	fmt.Println("err", err)
	if err != nil {
		fmt.Println("Ingresar nuevamente un numero valido")
		Excercise2()
	} else {
		for i := 1; i <= 10; i++ {
			fmt.Println(numero1, " x ", i, " = ", numero1*i)
		}

	}
	// if scanner.Scan() {
	// 	if err != nil{
	// 		continue
	// 	}else {
	// 		break
	// 	}

	// }
}
