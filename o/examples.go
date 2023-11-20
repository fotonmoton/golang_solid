package o

import "fmt"

func NotOpenClosed() {
	fmt.Printf("Result for 1 + 2 is: %d\n", NotExtendableCalculator("+", 1, 2))
	fmt.Printf("Result for 2 - 2 is: %d\n", NotExtendableCalculator("-", 2, 2))
	NotExtendableRide(&Car{}, &Train{})
	NewNotExtendableUser("Greg").PrintName()
}

func OpenClosed() {
	fmt.Printf("Result for 1 + 2 is: %d\n", ExtendableCalculator(Addition{}, 1, 2))
	fmt.Printf("Result for 2 - 2 is: %d\n", ExtendableCalculator(Subtraction{}, 2, 2))
	ExtendableRide(&Car{}, &Train{})
}
