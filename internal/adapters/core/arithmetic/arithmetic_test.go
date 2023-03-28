package arithmetic

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAddition(t *testing.T) {
	arith := NewArithmeticAdapter()

	var value1 int32 = 10
	var value2 int32 = 20

	expected := value1 + value2

	funcResult, err := arith.Addition(value1, value2)

	if err != nil {
		log.Printf("Failed to run func addition %v", err)
	}

	assert.EqualValues(t, expected, funcResult, "FAILED TestAddition(%d,%d) expected = %d , got %d", value1, value2, expected, funcResult)
}

func TestDivision(t *testing.T) {
	arith := NewArithmeticAdapter()

	var value1 int32 = 10
	var value2 int32 = 20

	expected := value1 / value2

	funcResult, err := arith.Division(value1, value2)

	if err != nil {
		log.Printf("Failed to run division %v", err)
	}

	assert.EqualValues(t, expected, funcResult, "FAILED TestDivision(%d,%d) expected = %d , got %d", value1, value2, expected, funcResult)

}
func TestMultiplication(t *testing.T) {
	arith := NewArithmeticAdapter()

	var value1 int32 = 10
	var value2 int32 = 20

	expected := value1 * value2

	funcResult, err := arith.Multiplication(value1, value2)

	if err != nil {
		log.Printf("Failed to run multiplication %v", err)
	}

	assert.EqualValues(t, expected, funcResult, "FAILED TestMultiplication%d,%d) expected = %d , got %d", value1, value2, expected, funcResult)

}
func TestSubtraction(t *testing.T) {
	arith := NewArithmeticAdapter()

	var value1 int32 = 10
	var value2 int32 = 20

	expected := value1 - value2

	funcResult, err := arith.Subtraction(value1, value2)

	if err != nil {
		log.Printf("Failed to run subtraction %v", err)
	}

	assert.EqualValues(t, expected, funcResult, "FAILED TestSubtraction(%d,%d) expected = %d , got %d", value1, value2, expected, funcResult)

}
