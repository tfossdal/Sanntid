package elevio

import (

)

type State int

const (
	Idle     State = 0
	Moving   State = 1
	DoorOpen State = 2
)
