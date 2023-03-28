package ports

type DbPort interface {
	CloseDbConnection()
	AddToHistory(answer int32, operation string) error
}

type MsgBrokerPort interface {
	PublishMessage(msg string) error
}
