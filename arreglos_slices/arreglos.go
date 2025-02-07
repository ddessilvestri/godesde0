package arreglos_slices

import "fmt"

var table [10]int = [10]int{10, 0, 59}
var matriz [20][30]int

func MuestroArreglos() {
	table[7] = 33
	table[2] = 54

	tabla2 := [100]string{"Pablo", "Juan"}
	fmt.Println(table)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)
}
