package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func calcPrice(width, height float64, chambers int, material string, hasSill bool) float64 {
	area := width * height

	priceTable := map[string]map[int]float64{
		"Дерев'яний": {
			1: 2.5,
			2: 3,
		},
		"Металевий": {
			1: 0.5,
			2: 1,
		},
		"Металопластиковий": {
			1: 1.5,
			2: 2,
		},
	}

	price := area * priceTable[material][chambers]

	if hasSill {
		price += 350
	}

	return price
}

func main() {
	a := app.New()
	w := a.NewWindow("Розрахунок склопакета")
	w.Resize(fyne.NewSize(500, 350))

	// Підписи та поля введення
	widthLabel := widget.NewLabel("Ширина (см):")
	widthEntry := widget.NewEntry()
	widthEntry.SetPlaceHolder("Введіть ширину склопакета в см")

	heightLabel := widget.NewLabel("Висота (см):")
	heightEntry := widget.NewEntry()
	heightEntry.SetPlaceHolder("Введіть висоту склопакета в см")

	chambersLabel := widget.NewLabel("Кількість камер:")
	chambers := widget.NewSelect([]string{"Однокамерний", "Двокамерний"}, func(string) {})
	chambers.SetSelected("Однокамерний")

	materialLabel := widget.NewLabel("Матеріал рами:")
	material := widget.NewSelect([]string{
		"Дерев'яний",
		"Металевий",
		"Металопластиковий",
	}, func(string) {})
	material.SetSelected("Дерев'яний")

	sill := widget.NewCheck("Додати підвіконня (350 грн)", nil)

	result := widget.NewLabel("")

	button := widget.NewButton("Розрахувати", func() {
		wVal, _ := strconv.ParseFloat(widthEntry.Text, 64)
		hVal, _ := strconv.ParseFloat(heightEntry.Text, 64)

		var ch int
		if chambers.Selected == "Однокамерний" {
			ch = 1
		} else {
			ch = 2
		}

		mat := material.Selected

		total := calcPrice(wVal, hVal, ch, mat, sill.Checked)
		result.SetText(fmt.Sprintf("Вартість: %.2f грн", total))
	})

	w.SetContent(container.NewVBox(
		widthLabel,
		widthEntry,
		heightLabel,
		heightEntry,
		chambersLabel,
		chambers,
		materialLabel,
		material,
		sill,
		button,
		result,
	))

	w.ShowAndRun()
}
