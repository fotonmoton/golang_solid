package s

import (
	"fmt"
	"time"
)

type Employee struct {
	Id        int
	FullName  string
	WorkedFor int
}

var hourlyRate = 10

var overtimeRate = hourlyRate * 2

var weeklyWorkingHours = 40

var employeeId = 0

func NewEmployee(name string, worked int) *Employee {
	employeeId++

	return &Employee{
		Id:        employeeId,
		FullName:  name,
		WorkedFor: worked,
	}
}

func (e *Employee) CalculatePay(worked int) int {

	// Switch lines to see difference:
	inHours := int(worked / int(time.Hour))
	// inHours := worked

	e.WorkedFor = worked

	if inHours > weeklyWorkingHours {
		return hourlyRate*weeklyWorkingHours + (inHours-weeklyWorkingHours)*overtimeRate
	}

	return inHours * hourlyRate
}

func (e *Employee) ReportHours() string {
	inHours := int(e.WorkedFor / int(time.Hour))
	return fmt.Sprintf("%s has worked for %d hours this week.", e.FullName, inHours)
}
