package s

import (
	"fmt"
	"time"
)

func NotResponsible() {

	greg := NewEmployee("Greg", 42)

	// Switch lines to see difference:
	fmt.Printf("Greg's weekly salary is: %d$\n", greg.CalculatePay(int(time.Duration(greg.WorkedFor)*time.Hour)))
	// fmt.Printf("Greg's weekly salary is: %d$\n", greg.CalculatePay(greg.WorkedFor))
	fmt.Println(greg.ReportHours())
}

func Responsible() {

	greg := NewEmployee("Greg", 42)

	// Switch lines to see difference:
	fmt.Printf("Greg's weekly salary is: %d$\n", CalculatePay(int(time.Duration(greg.WorkedFor)*time.Hour)))
	// fmt.Printf("Greg's weekly salary is: %d$\n", CalculatePay(greg.WorkedFor))
	fmt.Println(ReportHours(greg.FullName, time.Duration(greg.WorkedFor)*time.Hour))
}
