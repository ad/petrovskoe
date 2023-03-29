package app

import (
	"fmt"

	"github.com/ad/petrovskoe/internal/petrovskoe"
)

func Run() error {
	text := "Съешь ещё этих мягких французских булок да выпей чаю."

	fmt.Println("Оригинальный текст:")
	fmt.Println(text)

	petrovskyText := petrovskoe.ConvertToPetrovsky(text)

	fmt.Println("Текст в петровском правописании:")
	fmt.Println(petrovskyText)

	return nil
}
