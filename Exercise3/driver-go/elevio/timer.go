package elevio

import (
	"time"
)

var timerEndTime float32
var timerActive int

func Timer_start(duration float32) {
	timerEndTime = float32(time.Now().Unix()) + duration
	timerActive = 1
}

func Timer_stop() {
	timerActive = 0
}

func Timer_timedOut() int {
	if timerActive != 0 && float32(time.Now().Unix()) > timerEndTime {
		return 1
	}
	return 0
}
