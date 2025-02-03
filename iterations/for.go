package iterations

import (
	"fmt"
)

func Iterar() {
	i := 0
	for i < 10 {
		fmt.Println("Iterar", i)
		i++
	}

	for c := 0; c < 10; c++ {
		fmt.Println("Iterar in c++", c)

	}
}
