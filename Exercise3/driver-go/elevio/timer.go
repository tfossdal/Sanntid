package elevio

import (
	"fmt"
	"time"
)

var timerEndTime float64
var timerActive int

func Timer_start(duration float64) {
	timerEndTime = float64(time.Now().Unix()) + duration
	timerActive = 1
	fmt.Println("Timer started")
}

func Timer_stop() {
	timerActive = 0
}

func Timer_timedOut() int {
	if timerActive != 0 && float64(time.Now().Unix()) > timerEndTime {
		fmt.Println("Timed out")
		return 1
	}
	return 0
}

func CheckForTimeout() {
	for {
		if Timer_timedOut() != 0 {
			fmt.Print("hei")
			Timer_stop()
			Fsm_OnDoorTimeout()
			PrintState()
		}
	}
}
