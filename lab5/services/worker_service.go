package services

import (
	"lab5/models"
	"lab5/utils"
)

type WorkerService interface {
	ReadWorkers() []*models.Worker
	PrintWorker(*models.Worker)
	PrintWorkers([]*models.Worker)
	GetSalaryRange([]*models.Worker) (min float64, max float64)
}

type WorkerServiceImpl struct{}

func NewWorkerService() WorkerService {
	return &WorkerServiceImpl{}
}

func (s *WorkerServiceImpl) ReadWorkers() []*models.Worker {
	n := utils.ReadInt("Enter number of workers: ")
	return utils.ReadWorkers(n)
}

func (s *WorkerServiceImpl) PrintWorker(w *models.Worker) {
	utils.PrintWorker(w)
}

func (s *WorkerServiceImpl) PrintWorkers(workers []*models.Worker) {
	for _, w := range workers {
		s.PrintWorker(w)
	}
}

func (s *WorkerServiceImpl) GetSalaryRange(workers []*models.Worker) (min float64, max float64) {
	if len(workers) == 0 {
		return 0, 0
	}

	min = workers[0].Company.Salary
	max = workers[0].Company.Salary

	for _, w := range workers {
		sal := w.Company.Salary
		if sal < min {
			min = sal
		}
		if sal > max {
			max = sal
		}
	}

	return min, max
}
