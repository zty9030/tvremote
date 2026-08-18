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

// SendKey sends a normal key press: DOWN + UP.
func (s *Service) SendKey(key string) error {
    action, ok := keyMap[key]
    if !ok {
        return fmt.Errorf("unknown key: %s", key)
    }

    switch action.Type {
    case "keyevent":
        command := fmt.Sprintf(
            "input keyevent %d",
            action.Code,
        )

        return s.executor.Execute(command)

    case "sendevent":
        return s.sendEvent(action)

    default:
        return fmt.Errorf(
            "unknown action type: %s",
            action.Type,
        )
    }
}

// KeyDown sends only the key-down event.
func (s *Service) KeyDown(key string) error {
    action, ok := keyMap[key]
    if !ok {
        return fmt.Errorf("unknown key: %s", key)
    }

    if action.Type != "sendevent" {
        return fmt.Errorf(
            "key %s does not support key down",
            key,
        )
    }

    return s.sendEventDown(action)
}

// KeyUp sends only the key-up event.
func (s *Service) KeyUp(key string) error {
    action, ok := keyMap[key]
    if !ok {
        return fmt.Errorf("unknown key: %s", key)
    }
    if action.Type != "sendevent" {
        return fmt.Errorf(
            "key %s does not support key up",
            key,
        )
    }
    return s.sendEventUp(action)
}

// sendEvent sends a complete DOWN + SYN + UP + SYN sequence.
func (s *Service) sendEvent(action KeyAction) error {
    if err := s.sendEventDown(action); err != nil {
        return err
    }
    return s.sendEventUp(action)
}

// sendEventDown sends:
// EV_KEY DOWN
// SYN_REPORT
func (s *Service) sendEventDown(action KeyAction) error {
    command := fmt.Sprintf(
        "sendevent %s 1 %d 1 && sendevent %s 0 0 0",
        action.Device,
        action.Code,
        action.Device,
    )
    return s.executor.Execute(command)
}

// sendEventUp sends:
// EV_KEY UP
// SYN_REPORT
func (s *Service) sendEventUp(action KeyAction) error {
    command := fmt.Sprintf(
        "sendevent %s 1 %d 0 && sendevent %s 0 0 0",
        action.Device,
        action.Code,
        action.Device,
    )
    return s.executor.Execute(command)
}
