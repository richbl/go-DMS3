package dms3libs

// MotionDetector is the motion detector application run on the clients
var MotionDetector = structMotionDetector{
	command: "motion",
	state:   Stop,
}

// Command returns the motion detector application command
func (s structMotionDetector) Command() string {
	return s.command
}

// State returns the motion detector application state
func (s structMotionDetector) State() MotionDetectorState {
	return s.state
}

// SetState sets the motion detector application state
func (s *structMotionDetector) SetState(state MotionDetectorState) {
	s.state = state
}

// states of the motion detector application
const (
	Stop MotionDetectorState = iota
	Start
)

type structMotionDetector struct {
	command string
	state   MotionDetectorState
}

// MotionDetectorState defines the motion detector application state type
type MotionDetectorState int
