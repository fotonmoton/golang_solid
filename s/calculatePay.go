package s

import "time"

func CalculatePay(worked int) int {

	// Switch lines to see difference:
	inHours := int(worked / int(time.Hour))
	// inHours := worked

	if inHours > weeklyWorkingHours {
		return hourlyRate*weeklyWorkingHours + (inHours-weeklyWorkingHours)*overtimeRate
	}

	return inHours * hourlyRate
}
