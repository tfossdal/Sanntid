package elevio

func Requests_ShouldClearImmediately(e Elevator, btn_floor int, btn_type ButtonType) int {
	switch e.config.clearRequestVariant {
	case CV_ALL:
		if e.floor == btn_floor {
			return 1
		}
		return 0
	case CV_InDirn:
		if e.floor == btn_floor &&
			(e.dirn == MD_Up && btn_type == BT_HallUp) ||
			(e.dirn == MD_Down && btn_type == BT_HallDown) ||
			e.dirn == MD_Stop ||
			btn_type == BT_Cab {
			return 1
		}
		return 0
	default:
		return 0
	}
}

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

func Requests_shouldStop(e Elevator) {
	switch elevator.dirn {
	default:
		break
	}
}
