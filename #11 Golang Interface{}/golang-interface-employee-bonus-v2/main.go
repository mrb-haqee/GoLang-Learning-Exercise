package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

func (j Junior) GetBonus() float64 {
	JB := float64(j.BaseSalary)
	JW := float64(j.WorkingMonth)
	month := float64(12)
	return JB * JW / month
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

func (s Senior) GetBonus() float64 {
	SB := float64(s.BaseSalary)
	SW := float64(s.WorkingMonth)
	SP := float64(s.PerformanceRate)
	month := float64(12)
	return float64(2)*SB*SW/month + SP*SB
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (m Manager) GetBonus() float64 {
	MB := float64(m.BaseSalary)
	MW := float64(m.WorkingMonth)
	MP := float64(m.PerformanceRate)
	MBM := float64(m.BonusManagerRate)
	month := float64(12)
	return float64(2)*MB*MW/month + MP*MB + MBM*MB
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus() // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	var sum float64
	for _, Bonus := range employees {
		sum += Bonus.GetBonus()
	}
	return sum // TODO: replace this
}

func main() {
	var Bonus Employee
	var Bonus2 Employee
	var Bonus3 Employee

	Bonus = Junior{Name: "Junior 1", BaseSalary: 100000, WorkingMonth: 12}
	Bonus2 = Senior{Name: "rafli", BaseSalary: 10_000_000, WorkingMonth: 13, PerformanceRate: 2.5}
	Bonus3 = Manager{Name: "rafli", BaseSalary: 10_000_000, WorkingMonth: 20, PerformanceRate: 2.5, BonusManagerRate: 2}
	fmt.Println(EmployeeBonus(Bonus))
	fmt.Println(EmployeeBonus(Bonus2))
	fmt.Println(EmployeeBonus(Bonus3))

	// var THR []Employee
	// var THR2 []Employee
	// var THR3 []Employee

	THR := []Employee{
		Junior{Name: "Junior A", BaseSalary: 100000, WorkingMonth: 12}}
	THR2 := []Employee{
		Senior{Name: "Senior A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5},
		Senior{Name: "Senior B", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5}}
	THR3 := []Employee{
		Junior{Name: "Junior A", BaseSalary: 100000, WorkingMonth: 12},
		Senior{Name: "Senior A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5},
		Manager{Name: "Manager A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5, BonusManagerRate: 0.5}}

	fmt.Println(TotalEmployeeBonus(THR))
	fmt.Println(TotalEmployeeBonus(THR2))
	fmt.Println(TotalEmployeeBonus(THR3))

}
