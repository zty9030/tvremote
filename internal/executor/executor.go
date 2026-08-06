package executor


type AndroidExecutor interface {
	Execute(command string) error
}