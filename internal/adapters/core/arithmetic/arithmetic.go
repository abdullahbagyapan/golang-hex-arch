package arithmetic

type Adapter struct {
}

func NewArithmeticAdapter() *Adapter {
	return new(Adapter) // &Adapter{}
}

func (arith Adapter) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

func (arith Adapter) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

func (arith Adapter) Subtraction(a, b int32) (int32, error) {
	return a - b, nil
}

func (arith Adapter) Division(a, b int32) (int32, error) {
	return a / b, nil
}
