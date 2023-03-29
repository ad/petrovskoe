package petrovskoe

// Функция задает порядок правил для преобразования слова в петровском стиле.
func order(text string) string {
	text = Yat(text)
	text = Er(text)
	text = i(text)
	text = Fita(text)
	text = Izhitsa(text)

	return text
}
