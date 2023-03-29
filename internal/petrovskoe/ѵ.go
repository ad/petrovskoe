package petrovskoe

import "strings"

/*
* Когда писать ѵ (ижицу)?
*
* Почти никогда.
* Ижица сохранилась только в слове мѵро (миро — церковный елей) и в некоторых других церковных терминах:
* ѵподіаконъ, ѵпостась и др. Эта буква также греческого происхождения, соответствует греческой букве «ипсилон».
 */
func Izhitsa(text string) string {
	text = strings.ReplaceAll(text, "миро", "мѵро")
	text = strings.ReplaceAll(text, "иподіаконъ", "ѵподіаконъ")
	text = strings.ReplaceAll(text, "ипост", "ѵпост")

	return text
}
