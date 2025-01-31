package variables

import "fmt"

func ShowIntegers() {
	var commonint int
	var intof16Min int16
	intof16 := int16(10)
	intof32 := int32(10)
	intof64 := int64(64)

	fmt.Println("commonint =", commonint)
	fmt.Println("intof16Min =", intof16Min)
	fmt.Println("intof16 =", intof16)
	fmt.Println("intof32 =", intof32)
	fmt.Println("intof64 =", intof64)
}
