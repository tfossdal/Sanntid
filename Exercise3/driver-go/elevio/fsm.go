package elevio

import (
	"fmt"
)

var state State = Idle

func Fsm_OnRequestButtonPress(btn_Floor int, btn_type ButtonType) {
	switch state {
	case Idle:
		fmt.Print("test")
	}
}
