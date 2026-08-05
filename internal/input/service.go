package input

import (
	"fmt"

	"tvremote/internal/shell"
)

func SendKey(key string) (*shell.Result, error) {

	code, ok := keyMap[key]
	if !ok {
		return nil, fmt.Errorf("unknown key: %s", key)
	}

	cmd := fmt.Sprintf("input keyevent %d", code)

	return shell.Execute(cmd)
}