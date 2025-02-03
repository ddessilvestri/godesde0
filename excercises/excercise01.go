package excercises

import (
	"strconv"
)

func Excercise1(text string) (int64, string, error) {
	textToInt, err := strconv.ParseInt(text, 10, 64)
	// textToInt, err := strconv.Atoi(text)

	var TextToReturn string
	// switch {
	// case textToInt > 100:
	// 	TextToReturn = "Es Mayor a 100"
	// default:
	// 	TextToReturn = "Menor o igual a 100"
	// }
	if textToInt > 100 {
		TextToReturn = "Es Mayor a 100"

	} else {
		TextToReturn = "Menor o igual a 100"

	}
	// fmt.Println(err)

	return textToInt, TextToReturn, err
}
