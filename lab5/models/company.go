package models

type Company struct {
	Name     string
	Position string
	Salary   float64
}

func NewCompany(name string, pos string, sal float64) *Company {
	return &Company{name, pos, sal}
}

func (c *Company) SetName(n string)     { c.Name = n }
func (c *Company) SetPosition(p string) { c.Position = p }
func (c *Company) SetSalary(s float64)  { c.Salary = s }

func (c *Company) GetName() string     { return c.Name }
func (c *Company) GetPosition() string { return c.Position }
func (c *Company) GetSalary() float64  { return c.Salary }
