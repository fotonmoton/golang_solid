package o

func NotExtendableCalculator(operation string, firstOperand int, secondOperand int) int {
	if operation == "+" {
		return firstOperand + secondOperand
	}

	if operation == "-" {
		return firstOperand - secondOperand
	}

	return -1
}

type operation interface {
	Execute(int, int) int
}

type Addition struct{}

func (a Addition) Execute(firstOperand int, secondOperand int) int {
	return firstOperand + secondOperand
}

type Subtraction struct{}

func (s Subtraction) Execute(firstOperand int, secondOperand int) int {
	return firstOperand - secondOperand
}

func ExtendableCalculator(operation operation, firstOperand int, secondOperand int) int {
	return operation.Execute(firstOperand, secondOperand)
}
