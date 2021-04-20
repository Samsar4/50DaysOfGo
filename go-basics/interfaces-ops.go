package main

import (
	"fmt"
)

type Operation interface {
	Calc() int
}

// type Sum addressing 2 operands as integers
type Sum struct {
	op1, op2 int
}

// method calculate (sum)
func (s Sum) Calc() int {
	return s.op1 + s.op2
}

// string representation of values
func (s Sum) String() string {
	return fmt.Sprintf("%d + %d", s.op1, s.op2)
}

// type Subtract addressing 2 operands as integers
type Subtract struct {
	op1, op2 int
}

// method calculate (subtraction)
func (s Subtract) Calc() int {
	return s.op1 - s.op2
}

// string representation of values
func (s Subtract) String() string {
	return fmt.Sprintf("%d - %d", s.op1, s.op2)
}

func main() {
	ops := make([]Operation, 4)
	ops[0] = Sum{50, 150}
	ops[1] = Subtract{20, 19}
	ops[2] = Subtract{10, 50}
	ops[3] = Sum{5, 2}

	counter := 0

	for _, op := range ops {
		value := op.Calc()
		fmt.Printf("%v = %d\n", op, value)
		counter += value
	}
	fmt.Println("Total =", counter)

	// declaring variable type Operation
	var sum Operation
	sum = Sum{60000, 5535}
	fmt.Printf("Another operation: %v = %d\n", sum, sum.Calc())
}
