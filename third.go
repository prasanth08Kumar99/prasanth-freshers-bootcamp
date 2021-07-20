package main

type SalaryCaluculator interface {
	Caluculate() int
}

type Employee struct {
	Type string
	Basic int
	Duration int
}

func (employee *Employee) Caluculate() int {
	return employee.Basic * employee.Duration
}

func main() {
	fulltime := Employee{"Full-Time", 500, 28}
	contract := Employee{"Contract", 100, 28}
	freelancer := Employee{"Freelancer", 10, 40}

	salaries := []SalaryCaluculator{fulltime, contract, freelancer}

}
