package ports

import (
	"github.com/yasir2000/my-hexagonal-go/internal/ports"
)

//API ports provide access to package that contain application that defines tha interface

type APIPort interface {
	// method signature
	GetAddition(a, b int32) (int32, error)
	GetSubtraction(a, b int32) (int32, error)
	GetMultiplication(a, b int32) (int32, error)
	GetDivision(a, b int32) (int32, error)
}

type Adaptor struct {
	arith ports.ArithmaticPort
}

//API adaptor implements the Adaptor ports interface
func NewAdaptor(arith ports.ArithmaticPort) *Adaptor {
	//dependency injection
	return &Adaptor{arith: arith}
}

func (apia ArithmaticPort) GetAddition(a int32, b int32) (int32, err) {
	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia ArithmaticPort) GetSubtraction(a int32, b int32) (int32, err) {
	answer, err := apia.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia ArithmaticPort) GetMultiplication(a int32, b int32) (int32, err) {
	answer, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}
func (apia ArithmaticPort) GetDivision(a int32, b int32) (int32, err) {
	answer, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}
