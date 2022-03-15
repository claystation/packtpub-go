package payroll

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

func (m Manager) Pay() (string, float64) {
	fullName := m.Individual.FullName()
	return fullName, m.Salary + (m.Salary * m.CommissionRate)
}
