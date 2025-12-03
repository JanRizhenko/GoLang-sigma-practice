package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lab5/models"
)

func ReadString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text != "" {
			return text
		}
		fmt.Println("Input cannot be empty.")
	}
}

func ReadInt(prompt string) int {
	for {
		fmt.Print(prompt)
		var raw string
		_, err := fmt.Scan(&raw)
		if err != nil {
			return 0
		}
		if val, err := strconv.Atoi(raw); err == nil {
			return val
		}
		fmt.Println("Invalid integer. Try again.")
	}
}

func ReadFloat(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var raw string
		_, err := fmt.Scan(&raw)
		if err != nil {
			return 0
		}
		if val, err := strconv.ParseFloat(raw, 64); err == nil {
			return val
		}
		fmt.Println("Invalid number. Try again.")
	}
}

func ReadWorkers(n int) []*models.Worker {
	workers := make([]*models.Worker, 0, n)

	for i := 0; i < n; i++ {
		fmt.Printf("\n--- Worker %d ---\n", i+1)

		name := ReadString("Name: ")
		year := ReadInt("Start year: ")
		month := ReadInt("Start month: ")
		place := ReadString("Workplace: ")

		companyName := ReadString("Company name: ")
		position := ReadString("Position: ")
		salary := ReadFloat("Salary: ")

		company := models.NewCompany(companyName, position, salary)
		worker := models.NewWorker(name, year, month, place, *company)

		workers = append(workers, worker)
	}

	return workers
}
