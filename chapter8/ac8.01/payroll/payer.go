package payroll

type Payer interface {
	Pay() (string, float64)
}
