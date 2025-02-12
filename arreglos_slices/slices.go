package arreglos_slices

import "fmt"

var tabla []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 5, 4, 89, 6, 5, 4, 89}

func MuestroSlice() {
	fmt.Println(tabla)
	fmt.Println(arreglo)

	procion := arreglo[3:]   // Slice creado desde un vector desde la posicion 3
	procion2 := arreglo[:5]  // Slice creado desde la posicion 0 a la 5
	procion3 := arreglo[6:7] // Slice creado desde la posicion 6 a la 7
	fmt.Println(procion)
	fmt.Println(procion2)
	fmt.Println(procion3)

}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
	nums := make([]int, 0)
	for i := range 100 {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
