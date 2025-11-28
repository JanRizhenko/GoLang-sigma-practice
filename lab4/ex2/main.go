package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func calcTourPrice(days int, country, season string, guide bool, luxury bool) float64 {
	priceTable := map[string]map[string]float64{
		"Болгарія": {
			"літо": 100,
			"зима": 150,
		},
		"Німеччина": {
			"літо": 160,
			"зима": 200,
		},
		"Польща": {
			"літо": 120,
			"зима": 180,
		},
	}

	price := float64(days) * priceTable[country][season]

	if guide {
		price += float64(days) * 50
	}

	if luxury {
		price *= 1.2
	}

	return price
}

func main() {
	a := app.New()
	w := a.NewWindow("Розрахунок туру")
	w.Resize(fyne.NewSize(500, 350))

	daysLabel := widget.NewLabel("Кількість днів:")
	daysEntry := widget.NewEntry()
	daysEntry.SetPlaceHolder("Введіть кількість днів подорожі")

	countryLabel := widget.NewLabel("Країна подорожі:")
	country := widget.NewSelect([]string{"Болгарія", "Німеччина", "Польща"}, func(string) {})
	country.SetSelected("Болгарія")

	seasonLabel := widget.NewLabel("Сезон поїздки:")
	season := widget.NewSelect([]string{"літо", "зима"}, func(string) {})
	season.SetSelected("літо")

	guide := widget.NewCheck("Індивідуальний гід (+50$/день)", nil)
	luxury := widget.NewCheck("Номер люкс (+20%)", nil)

	result := widget.NewLabel("")

	button := widget.NewButton("Розрахувати", func() {
		days, _ := strconv.Atoi(daysEntry.Text)
		c := country.Selected
		s := season.Selected
		g := guide.Checked
		l := luxury.Checked

		total := calcTourPrice(days, c, s, g, l)
		result.SetText(fmt.Sprintf("Вартість: %.2f $", total))
	})

	w.SetContent(container.NewVBox(
		daysLabel,
		daysEntry,
		countryLabel,
		country,
		seasonLabel,
		season,
		guide,
		luxury,
		button,
		result,
	))

	w.ShowAndRun()
}
