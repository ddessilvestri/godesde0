package excercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error

func Excercise2() string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
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
			text += fmt.Sprintf("%d x %d = %d \n", numero1, i, numero1*i)
		}

	}
	return text
	// if scanner.Scan() {
	// 	if err != nil{
	// 		continue
	// 	}else {
	// 		break
	// 	}

	// }
}
