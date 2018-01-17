package main

import "fmt"

type Employee struct {
	Fname,
	Ename string
	Pay float64
}

func (emp *Employee) GiveRaise() {
	emp.Pay *= 1.04
}

type Boss struct {
	Employee
}

func (boss *Boss) GiveRaise() {
	boss.Pay *= 1.01
}

type Developer struct {
	Employee
	Language string
}

func (dev *Developer) GiveRaise() {
	dev.Pay *= 1.10
}

func main() {
	emp1 := Employee{"Anna", "Jones", 34000}
	fmt.Println(emp1)
	emp1.GiveRaise()
	fmt.Println(emp1)

	dev1 := Developer{
		Employee{
			"Steve",
			"Smith",
			32000},
		"Go"}

	fmt.Printf("%#v \n", dev1)
	dev1.Employee.GiveRaise()
	fmt.Println(dev1.Fname)
}
