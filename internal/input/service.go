package input

import (
	"fmt"

	"tvremote/internal/executor"
)

type Service struct {
	executor executor.AndroidExecutor
}

func NewService(e executor.AndroidExecutor) *Service {

	return &Service{
		executor: e,
	}

}

func (s *Service) SendKey(key string) error {

	code, ok := keyMap[key]

	if !ok {

		return fmt.Errorf("unknown key: %s", key)

	}

	command := fmt.Sprintf(
		"input keyevent %d",
		code,
	)

	return s.executor.Execute(command)

}