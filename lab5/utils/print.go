package utils

import (
	"fmt"

	"lab5/models"
)

func PrintWorker(w *models.Worker) {
	fmt.Println("------------ Worker ------------")
	fmt.Println("Name:", w.Name)
	fmt.Printf("Start: %d-%02d\n", w.StartYear, w.StartMonth)
	fmt.Println("Workplace:", w.WorkPlace)
	fmt.Println("Position:", w.Company.Position)
	fmt.Printf("Salary: %.2f\n", w.Company.Salary)
}

func PrintWorkers(workers []*models.Worker) {
	for _, w := range workers {
		PrintWorker(w)
	}
}
