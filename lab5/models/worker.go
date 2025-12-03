package models

import "time"

type Worker struct {
	Name       string
	StartYear  int
	StartMonth int
	WorkPlace  string
	Company    Company
}

func NewWorker(name string, year, month int, place string, company Company) *Worker {
	return &Worker{
		Name:       name,
		StartYear:  year,
		StartMonth: month,
		WorkPlace:  place,
		Company:    company,
	}
}

func (w *Worker) SetName(name string)       { w.Name = name }
func (w *Worker) SetStartYear(year int)     { w.StartYear = year }
func (w *Worker) SetStartMonth(month int)   { w.StartMonth = month }
func (w *Worker) SetWorkPlace(place string) { w.WorkPlace = place }
func (w *Worker) SetCompany(c Company)      { w.Company = c }

func (w *Worker) GetName() string      { return w.Name }
func (w *Worker) GetStartYear() int    { return w.StartYear }
func (w *Worker) GetStartMonth() int   { return w.StartMonth }
func (w *Worker) GetWorkPlace() string { return w.WorkPlace }
func (w *Worker) GetCompany() Company  { return w.Company }

func (w *Worker) GetPosition() string {
	return w.Company.Position
}

func (w *Worker) GetWorkExperience() int {
	now := time.Now()
	years := now.Year() - w.StartYear
	months := int(now.Month()) - w.StartMonth
	return years*12 + months
}

func (w *Worker) GetTotalEarnings() float64 {
	return float64(w.GetWorkExperience()) * w.Company.Salary
}
