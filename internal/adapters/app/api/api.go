package api

import (
	"fmt"
	"hexarch/internal/ports"
	"log"
	"time"
)

type Adapter struct {
	arith     ports.ArithmeticPort
	db        ports.DbPort
	msgBroker ports.MsgBrokerPort
}

func NewAPIAdapter(db ports.DbPort, arith ports.ArithmeticPort, msgBroker ports.MsgBrokerPort) *Adapter {
	return &Adapter{db: db, arith: arith, msgBroker: msgBroker}
}

func (apiA Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apiA.arith.Addition(a, b)

	if err != nil {
		log.Printf("Error get addition %s", err)
	}

	err = apiA.db.AddToHistory(answer, "addition")

	if err != nil {
		log.Printf("Error to save db %s", err)
	}

	msg := fmt.Sprintf("Addition(%d,%d) used at %v", a, b, time.Now())

	err = apiA.msgBroker.PublishMessage(msg)

	if err != nil {
		log.Printf("Error to publish message %s", err)
	}

	return answer, nil

}
func (apiA Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apiA.arith.Multiplication(a, b)

	if err != nil {
		log.Printf("Error get multiplication %s", err)
	}

	err = apiA.db.AddToHistory(answer, "multiplication")

	if err != nil {
		log.Printf("Error to save db %s", err)
	}
	msg := fmt.Sprintf("Multiplication(%d,%d) used at %v", a, b, time.Now())

	err = apiA.msgBroker.PublishMessage(msg)

	if err != nil {
		log.Printf("Error to publish message %s", err)
	}

	return answer, nil
}
func (apiA Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apiA.arith.Division(a, b)

	if err != nil {
		log.Printf("Error get division %s", err)
	}
	err = apiA.db.AddToHistory(answer, "division")

	if err != nil {
		log.Printf("Error to save db %s", err)
	}
	msg := fmt.Sprintf("Division(%d,%d) used at %v", a, b, time.Now())

	err = apiA.msgBroker.PublishMessage(msg)

	if err != nil {
		log.Printf("Error to publish message %s", err)
	}
	return answer, nil
}
func (apiA Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apiA.arith.Subtraction(a, b)

	if err != nil {
		log.Printf("Error get subtraction %s", err)
	}
	err = apiA.db.AddToHistory(answer, "subtraction")

	if err != nil {
		log.Printf("Error to save db %s", err)
	}
	msg := fmt.Sprintf("Subtraction(%d,%d) used at %v", a, b, time.Now())

	err = apiA.msgBroker.PublishMessage(msg)

	if err != nil {
		log.Printf("Error to publish message %s", err)
	}
	return answer, nil
}
