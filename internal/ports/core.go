package ports

type ArithmeticPort interface {
	Multiplication(a, b int32) (int32, error)
	Addition(a, b int32) (int32, error)
	Subtraction(a, b int32) (int32, error)
	Division(a, b int32) (int32, error)
}
