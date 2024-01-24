package elevio

var elevator Elevator = Elevator{-1, MD_Stop, [_numFloors][_numButtons]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, Idle, Config{CV_ALL, 3.0}}

func Fsm_OnRequestButtonPress(btn_Floor int, btn_type ButtonType) {
	switch elevator.state {
	case DoorOpen:
		if Requests_ShouldClearImmediately(elevator, btn_Floor, btn_type) != 0 {
			Timer_start(elevator.config.doorOpenDuration_s)
		} else {
			elevator.requests[btn_Floor][btn_type] = 1
		}
	case Moving:
		elevator.requests[btn_Floor][btn_type] = 1
	case Idle:
		elevator.requests[btn_Floor][btn_type] = 1
		var pair DirnBehaviourPair = Requests_chooseDirection(elevator)
		elevator.dirn = pair.dirn
		elevator.state = pair.state
		switch pair.state {
		case DoorOpen:
			SetDoorOpenLamp(true)
			Timer_start(elevator.config.doorOpenDuration_s)
			//elevator = Request_clearAtCurrentFloor(elevator)
		case Moving:
			SetMotorDirection(elevator.dirn)
		case Idle:
			break
		}
	}
}

func Fsm_OnFloorArrival(newFloor int) {
	elevator.floor = newFloor

	SetFloorIndicator(newFloor)

	switch elevator.state {
	case Moving:
		if 1 != 0 {
			return
		}
	}
}
