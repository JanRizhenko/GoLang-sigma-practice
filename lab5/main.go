package main

import (
	"fmt"
	"lab5/services"
)

func main() {
	var workerService services.WorkerService
	workerService = services.NewWorkerService()

	workers := workerService.ReadWorkers()

	fmt.Println("\n=== Workers List ===")
	workerService.PrintWorkers(workers)

	minSal, maxSal := workerService.GetSalaryRange(workers)
	fmt.Printf("\nMinimum salary: %.2f\n", minSal)
	fmt.Printf("Maximum salary: %.2f\n", maxSal)
}
