package main

/*
#include <stdbool.h>
double calc(double width, double height, int chambers, int material, bool sill);
*/
import "C"

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Розрахунок склопакета")
	w.Resize(fyne.NewSize(500, 350))

	widthLabel := widget.NewLabel("Ширина (см):")
	widthEntry := widget.NewEntry()
	widthEntry.SetPlaceHolder("Введіть ширину в см")

	heightLabel := widget.NewLabel("Висота (см):")
	heightEntry := widget.NewEntry()
	heightEntry.SetPlaceHolder("Введіть висоту в см")

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

		var mat int
		switch material.Selected {
		case "Дерев'яний":
			mat = 0
		case "Металевий":
			mat = 1
		case "Металопластиковий":
			mat = 2
		}

		total := C.calc(
			C.double(wVal),
			C.double(hVal),
			C.int(ch),
			C.int(mat),
			C.bool(sill.Checked),
		)

		result.SetText(fmt.Sprintf("Вартість: %.2f грн", float64(total)))
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
