package protocol

import "fmt"

type ServoControl int

const (
	ServoDisabled ServoControl = 0
	ServoNormal   ServoControl = 1
	ServoGantry   ServoControl = 8
)

// Sets the maximum velocity for jogging moves (units/ms)
func (p *PowerPMAC) SetJogSpeed(id int, speed float64) error {
	command := fmt.Sprintf("Motor[%d].JogSpeed=%f", id, speed)
	_, err := p.RequestFloat(command)
	return err
}

// Gets the maximum velocity for jogging moves (units/ms)
func (p *PowerPMAC) GetJogSpeed(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].JogSpeed", id)
	return p.RequestFloat(command)
}

// Sets the acceleration/deceleration for jogging moves (units/ms²)
func (p *PowerPMAC) SetJogAcceleration(id int, speed float64) error {
	command := fmt.Sprintf("Motor[%d].JogTa=%f", id, speed)
	_, err := p.RequestFloat(command)
	return err
}

// Gets the acceleration/deceleration for jogging moves (units/ms²)
func (p *PowerPMAC) GetJogAcceleration(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].JogTa", id)
	return p.RequestFloat(command)
}

// Sets the jerk for jogging moves (units/ms³)
func (p *PowerPMAC) SetJogJerk(id int, speed float64) error {
	command := fmt.Sprintf("Motor[%d].JogTs=%f", id, speed)
	_, err := p.RequestFloat(command)
	return err
}

// Gets the jerk for jogging moves (units/ms³)
func (p *PowerPMAC) GetJogJerk(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].JogTs", id)
	return p.RequestFloat(command)
}

func (p *PowerPMAC) SetHomeVelocity(id int, value float64) error {
	command := fmt.Sprintf("Motor[%d].HomeVel=%f", id, value)
	_, err := p.RequestFloat(command)
	return err
}

// Gets the magnitude and direction of the homing search move (units/ms)
func (p *PowerPMAC) GetHomeVelocity(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].HomeVel", id)
	return p.RequestFloat(command)
}

// Sets the distance from the trigger point to the final zero position (motor units)
func (p *PowerPMAC) SetHomeOffset(id int, value float64) error {
	command := fmt.Sprintf("Motor[%d].HomeOffset=%f", id, value)
	_, err := p.RequestFloat(command)
	return err
}

// Gets the distance from the trigger point to the final zero position (motor units)
func (p *PowerPMAC) GetHomeOffset(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].HomeOffset", id)
	return p.RequestFloat(command)
}

// Sets the global hardware-level velocity limit for all move types (units/ms)
func (p *PowerPMAC) SetMaximumSpeed(id int, value float64) error {
	command := fmt.Sprintf("Motor[%d].MaxSpeed=%f", id, value)
	_, err := p.RequestFloat(command)
	return err
}

// Gets the global hardware-level velocity limit for all move types (units/ms)
func (p *PowerPMAC) GetMaximumSpeed(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].MaxSpeed", id)
	return p.RequestFloat(command)
}

// Sets the positive software overtravel limit (motor units)
func (p *PowerPMAC) SetMaximumPosition(id int, value float64) error {
	command := fmt.Sprintf("Motor[%d].MaxPos=%f", id, value)
	_, err := p.RequestFloat(command)
	return err
}

// Gets the positive software overtravel limit (motor units)
func (p *PowerPMAC) GetMaximumPosition(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].MaxPos", id)
	return p.RequestFloat(command)
}

// Sets the negative software overtravel limit (motor units)
func (p *PowerPMAC) SetMinimumPosition(id int, value float64) error {
	command := fmt.Sprintf("Motor[%d].MinPos=%f", id, value)
	_, err := p.RequestFloat(command)
	return err
}

// Gets the negative software overtravel limit (motor units)
func (p *PowerPMAC) GetMininumPosition(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].MinPos", id)
	return p.RequestFloat(command)
}

// Sets the following error magnitude that causes a fatal shutdown (motor units)
func (p *PowerPMAC) SetFatalFollowingErrorLimit(id int, value float64) error {
	command := fmt.Sprintf("Motor[%d].FatalFeLimit=%f", id, value)
	_, err := p.RequestFloat(command)
	return err
}

// Gets the following error magnitude that causes a fatal shutdown (motor units)
func (p *PowerPMAC) GetFatalFollowingErrorLimit(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].FatalFeLimit", id)
	return p.RequestFloat(command)
}

// Sets the threshold for triggering a warning status bit (motor units)
func (p *PowerPMAC) SetWarningFollowingErrorLimit(id int, value float64) error {
	command := fmt.Sprintf("Motor[%d].WarnFeLimit=%f", id, value)
	_, err := p.RequestFloat(command)
	return err
}

// Gets the threshold for triggering a warning status bit (motor units)
func (p *PowerPMAC) GetWarningFollowingErrorLimit(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].WarnFeLimit", id)
	return p.RequestFloat(command)
}

// Sets the output/current clamping limit (range ±32,768)
func (p *PowerPMAC) SetMaximumOutput(id int, value float64) error {
	command := fmt.Sprintf("Motor[%d].MaxDac=%f", id, value)
	_, err := p.RequestFloat(command)
	return err
}

// Gets the output/current clamping limit (range ±32,768)
func (p *PowerPMAC) GetMaximumOutput(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].MaxDac", id)
	return p.RequestFloat(command)
}

// Sets the servo control state
func (p *PowerPMAC) SetServoControl(id int, mode ServoControl) error {
	command := fmt.Sprintf("Motor[%d].ServoCtrl=%d", id, mode)
	_, err := p.Request(command)
	return err
}

// Gets the servo control state
func (p *PowerPMAC) GetServoControl(id int) (ServoControl, error) {
	command := fmt.Sprintf("Motor[%d].ServoCtrl", id)
	value, err := p.RequestInt(command)
	return ServoControl(value), err
}

// Sets the leader motor used as reference to follow in gantry mode
func (p *PowerPMAC) SetLeaderMotor(id int, leader int) error {
	command := fmt.Sprintf("Motor[%d].CmdMotor=%d", id, leader)
	_, err := p.Request(command)
	return err
}

// Gets the leader motor used as reference to follow in gantry mode
func (p *PowerPMAC) GetLeaderMotor(id int) (int, error) {
	command := fmt.Sprintf("Motor[%d].CmdMotor", id)
	return p.RequestInt(command)
}

// Sets the error range for "in position" status (motor units)
func (p *PowerPMAC) SetInPositionBand(id int, units float64) error {
	command := fmt.Sprintf("Motor[%d].InPosBand=%f", id, units)
	_, err := p.Request(command)
	return err
}

// Gets the error range for "in position" status (motor units)
func (p *PowerPMAC) GetInPositionBand(id int) (float64, error) {
	command := fmt.Sprintf("#%dMotor[%d].InPosBand", id, id)
	return p.RequestFloat(command)
}

// Maps a motor to a coordinate system axis
func (p *PowerPMAC) SetMotorToAxis(coord int, id int, axis string) error {
	command := fmt.Sprintf("&%d #%d->%s", coord, id, axis)
	_, err := p.Request(command)
	return err
}
