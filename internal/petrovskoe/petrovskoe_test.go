package petrovskoe

import "testing"

func TestConvertToPetrovsky(t *testing.T) {
	type args struct {
		text string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test0",
			args: args{
				text: "Даниил Апатин",
			},
			want: "Данiилъ Апатинъ",
		},
		{
			name: "test1",
			args: args{
				text: "столъ, телефон, Санкт-Петербург!",
			},
			want: "столъ, телефонъ, Санктъ-Петербургъ!",
		},
		{
			name: "test2",
			args: args{
				text: "текст как дела?",
			},
			want: "текстъ какъ дела?",
		},
		{
			name: "test3",
			args: args{
				text: "текстs",
			},
			want: "текстs",
		},
		{
			name: "test4",
			args: args{
				text: "Санкт-Петербург",
			},
			want: "Санктъ-Петербургъ",
		},
		{
			name: "test5",
			args: args{
				text: "мяч, уж замуж невтерпеж",
			},
			want: "мячъ, ужъ замужъ невтерпежъ",
		},
		{
			name: "test6",
			args: args{
				text: "миро мир труд май",
			},
			want: "мѵро міръ трудъ май",
		},
		{
			name: "test7",
			args: args{
				text: "линия, другие, приѣхалъ, синий",
			},
			want: "линiя, другiе, прiѣхалъ, синiй",
		},
		{
			name: "test8",
			args: args{
				text: `Белый, блѣдный, бѣдный бѣсъ
Убѣжалъ голодный въ лѣсъ.
Лѣшимъ по лѣсу онъ бѣгалъ,
Рѣдькой съ хрѣномъ пообѣдалъ
И за горькій тотъ обѣдъ
Далъ обѣтъ надѣлать бѣдъ.

Вѣдай, братъ, что клѣть и клѣтка,
Рѣшето, рѣшетка, сѣтка,
Вѣжа и желѣзо съ ять —
Такъ и надобно писать.

Наши вѣки и рѣсницы
Защищаютъ глазъ зѣницы,
Вѣки жмуритъ цѣлый вѣкъ
Ночью каждый человѣкъ…

Вѣтеръ вѣтки поломалъ,
Нѣмецъ вѣники связалъ,
Свѣсилъ вѣрно при промѣнѣ,
За двѣ гривны продалъ въ Вѣнѣ.

Днѣпръ и Днѣстръ, какъ всѣмъ извѣстно,
Двѣ рѣки въ сосѣдствѣ тѣсномъ,
Дѣлитъ области ихъ Бугъ,
Рѣжетъ съ сѣвера на югъ.

Кто тамъ гнѣвно свирѣпѣетъ?
Крѣпко сѣтовать такъ смѣетъ?
Надо мирно споръ рѣшить
И другъ друга убѣдить…

Птичьи гнѣзда грѣхъ зорить,
Грѣхъ напрасно хлѣбъ сорить,
Надъ калѣкой грѣхъ смѣяться,
Надъ увѣчнымъ издѣваться…`,
			},
			want: `Бѣлый, блѣдный, бѣдный бѣсъ
Убѣжалъ голодный въ лѣсъ.
Лѣшимъ по лѣсу онъ бѣгалъ,
Рѣдькой съ хрѣномъ пообѣдалъ
И за горькій тотъ обѣдъ
Далъ обѣтъ надѣлать бѣдъ.

Вѣдай, братъ, что клѣть и клѣтка,
Рѣшето, рѣшетка, сѣтка,
Вѣжа и желѣзо съ ять —
Такъ и надобно писать.

Наши вѣки и рѣсницы
Защищаютъ глазъ зѣницы,
Вѣки жмуритъ цѣлый вѣкъ
Ночью каждый человѣкъ…

Вѣтеръ вѣтки поломалъ,
Нѣмецъ вѣники связалъ,
Свѣсилъ вѣрно при промѣнѣ,
За двѣ гривны продалъ въ Вѣнѣ.

Днѣпръ и Днѣстръ, какъ всѣмъ извѣстно,
Двѣ рѣки въ сосѣдствѣ тѣсномъ,
Дѣлитъ области ихъ Бугъ,
Рѣжетъ съ сѣвера на югъ.

Кто тамъ гнѣвно свирѣпѣетъ?
Крѣпко сѣтовать такъ смѣетъ?
Надо мирно споръ рѣшить
И другъ друга убѣдить…

Птичьи гнѣзда грѣхъ зорить,
Грѣхъ напрасно хлѣбъ сорить,
Надъ калѣкой грѣхъ смѣяться,
Надъ увѣчнымъ издѣваться…`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToPetrovsky(tt.args.text); got != tt.want {
				t.Errorf("ConvertToPetrovsky() = %v, want %v", got, tt.want)
			}
		})
	}
}
