package main

import (
	"hexarch/internal/adapters/app/api"
	"hexarch/internal/adapters/core/arithmetic"
	"hexarch/internal/adapters/framework/left/grpc"
	"hexarch/internal/adapters/framework/right/db"
	"hexarch/internal/adapters/framework/right/rabbitmq"
	"hexarch/internal/ports"
)

func main() {

	// ports
	var dbAdapter ports.DbPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort
	var core ports.ArithmeticPort
	var msgBroker ports.MsgBrokerPort

	// database
	dbAdapter = db.NewAdapter()
	// message broker
	msgBroker = rabbitmq.NewMsgBrokerAdapter()
	// core
	core = arithmetic.NewArithmeticAdapter()
	// application
	appAdapter = api.NewAPIAdapter(dbAdapter, core, msgBroker)
	// grpc
	grpcAdapter = grpc.NewGRPCAdapter(appAdapter)

	grpcAdapter.Run()

}
