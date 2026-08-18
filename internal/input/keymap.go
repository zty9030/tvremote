package input

type KeyAction struct {
    Type   string
    Code   int
    Device string
}

var keyMap = map[string]KeyAction{
    "UP": {
        Type: "keyevent",
        Code: 19,
    },
    "DOWN": {
        Type: "keyevent",
        Code: 20,
    },
    "LEFT": {
        Type: "keyevent",
        Code: 21,
    },
    "RIGHT": {
        Type: "keyevent",
        Code: 22,
    },
    "OK": {
        Type: "keyevent",
        Code: 23,
    },
    "BACK": {
        Type: "keyevent",
        Code: 4,
    },
    "HOME": {
        Type: "keyevent",
        Code: 3,
    },
    "MENU": {
        Type: "keyevent",
        Code: 82,
    },
    "VOL_UP": {
        Type:   "sendevent",
        Code:   30,
        Device: "/dev/input/event0",
    },
    "VOL_DOWN": {
        Type:   "sendevent",
        Code:   33,
        Device: "/dev/input/event0",
    },
    "POWER": {
        Type:   "sendevent",
        Code:   1, // KEY_ESC
        Device: "/dev/input/event0",
    },
    "SOURCE": {
        Type:   "sendevent",
        Code:   76,
        Device: "/dev/input/event0",
    },
    "MUTE": {
        Type: "keyevent",
        Code: 164,
    },
}
