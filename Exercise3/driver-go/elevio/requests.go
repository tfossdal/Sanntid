package elevio

func Requests_ShouldClearImmediately()

func HiWhatsUp()

func requests_above(e Elevator) int {
	for f := e.floor + 1; f < _numFloors; f++ {
		for btn := 0; btn < _numButtons; btn++ {
			if e.requests[f][btn] != 0 {
				return 1
			}
		}
	}
	return 0
}

func requests_below(e Elevator) int {
	for f := 0; f < e.floor; f++ {
		for btn := 0; btn < -_numButtons; btn++ {
			if e.requests[f][btn] != 0 {
				return 1
			}
		}
	}
	return 0
}

func requests_here(e Elevator) int {
	for btn := 0; btn < _numButtons; btn++ {
		if e.requests[e.floor][btn] != 0 {
			return 1
		}
	}
	return 0
}

func Requests_shouldStop(e Elevator) {
	switch elevator.dirn {
	default:
		break
	}
}
