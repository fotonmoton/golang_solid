package s

import (
	"fmt"
	"time"
)

func ReportHours(name string, worked time.Duration) string {
	inHours := int(worked / time.Hour)
	return fmt.Sprintf("%s has worked for %d hours this week.", name, inHours)
}
