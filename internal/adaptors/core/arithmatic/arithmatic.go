package arithamtic

type Adaptor struct {
}

//instance of Adaptor struct, instances mostly implemented with pointer type and returns
func NewAdaptor() *Adaptor {
	return &Adaptor{}
}

//implementation of interface method Addition
func (arith Adaptor) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

//implementation of interface method Subtraction
func (arith Adaptor) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

//implementation of interface method Multiplication
func (arith Adaptor) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

//implementation of interface method Division
func (arith Adaptor) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
