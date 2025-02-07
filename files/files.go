package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ddessilvestri/godesde0/excercises"
)

var fileName string = "./files/txt/tabla.txt"

func SaveTable() {
	var text string = excercises.Excercise2()
	archive, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	defer archive.Close() // This guarantees the file is closed when the function exits

	fmt.Fprintln(archive, text)

}

func AppendTableToFile() {
	var text string = excercises.Excercise2()
	if !Append(fileName, text) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(faile string, text string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	defer arch.Close()

	_, err = arch.WriteString(text)
	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
	}

	return true
}

func ReadFile() {
	archive, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo" + err.Error())
		return
	}

	fmt.Println(string(archive))

}

func ReadFile2() {
	archive, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo" + err.Error())
		return
	}
	defer archive.Close()
	scanner := bufio.NewScanner(archive)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}

}
