package contract

type CommandHandlerInterface interface {
	Handle(command interface{}) (interface{}, error)
	GetDtoOutputChannel() chan interface{}
}
