package functions

import "fmt" 

func table(valor int) func() int {
	numero :=valor 
	sequence:=0
	return func() int{
		sequence++
		return numero *sequence
	}
}

func LlamarClosure() {
	tabladel:=2 
	MiTabla :=table(tabladel) 
	for i:=1; i<=10; i++{
		fmt.Println(MiTabla())
	}
}