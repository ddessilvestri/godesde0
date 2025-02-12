package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires,Rio Negro"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 39,
		"Boca":        39,
		"Aldosivi":    119,
	}
	fmt.Println(campeonato)
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s,tiene un puntaje de %d \n", equipo, puntaje)
	}
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Barcelona"]
	fmt.Printf("El puntaje es %d, y el equipo existe = %t", puntaje, existe)

}
