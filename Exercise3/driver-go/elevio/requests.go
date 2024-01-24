package elevio

type DirnBehaviourPair struct {
	dirn  MotorDirection
	state State
}

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

func requests_here(e Elevator) int {
	for btn := 0; btn < _numButtons; btn++ {
		if e.requests[e.floor][btn] != 0 {
			return 1
		}
	}
	return 0
}

func Requests_chooseDirection(e Elevator) DirnBehaviourPair {
	switch e.dirn {
	case MD_Up:
		if requests_above(e) != 0 {
			return DirnBehaviourPair{MD_Up, Moving}
		}
		if requests_here(e) != 0 {
			return DirnBehaviourPair{MD_Down, DoorOpen}
		}
		if requests_below(e) != 0 {
			return DirnBehaviourPair{MD_Down, Moving}
		}
		return DirnBehaviourPair{MD_Stop, Idle}
	case MD_Down:
		if requests_below(e) != 0 {
			return DirnBehaviourPair{MD_Down, Moving}
		}
		if requests_here(e) != 0 {
			return DirnBehaviourPair{MD_Up, DoorOpen}
		}
		if requests_above(e) != 0 {
			return DirnBehaviourPair{MD_Up, Moving}
		}
		return DirnBehaviourPair{MD_Stop, Idle}
	case MD_Stop:
		if requests_here(e) != 0 {
			return DirnBehaviourPair{MD_Stop, DoorOpen}
		}
		if requests_above(e) != 0 {
			return DirnBehaviourPair{MD_Up, Moving}
		}
		if requests_below(e) != 0 {
			return DirnBehaviourPair{MD_Down, Moving}
		}
		return DirnBehaviourPair{MD_Stop, Idle}
	default:
		return DirnBehaviourPair{MD_Stop, Idle}
	}
}

func Requests_shouldStop(e Elevator) int {
	switch elevator.dirn {
	case MD_Down:
		if e.requests[e.floor][BT_HallDown]+e.requests[e.floor][BT_HallDown]-(requests_above(e)-1) != 0 {
			return 1
		}
		return 0
	case MD_Up:
		if e.requests[e.floor][BT_HallDown]+e.requests[e.floor][BT_HallUp]-(requests_above(e)-1) != 0 {
			return 1
		}
		return 0
	default:
		return 1
	}
}
