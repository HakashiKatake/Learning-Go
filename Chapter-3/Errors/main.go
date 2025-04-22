package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(4, 0)

	if err != nil {
		fmt.Println(err)
		// Do something with the error
		return
	}

	fmt.Println(result)
	// Use the result
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func Dividee(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}

	return a / b, nil
}

var ErrDivideByZero = errors.New("cannot divide by zero")

func DivideE(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	return a / b, nil
}

type DivisionError struct {
	Code int
	Msg  string
}

func (d DivisionError) Error() string {
	return fmt.Sprintf("code %d: %s", d.Code, d.Msg)
}

func DivideA(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionError{
			Code: 2000,
			Msg:  "cannot divide by zero",
		}
	}

	return a / b, nil
}
