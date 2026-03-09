package protocol

import "fmt"

// Gets the net actual position of the motor, including compensation
// and backslash
func (p *PowerPMAC) GetActualPosition(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].ActPos", id)
	return p.RequestFloat(command)
}

// Gets the net desired position
func (p *PowerPMAC) GetDesiredPosition(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].DesPos", id)
	return p.RequestFloat(command)
}

// Gets the home offset position
func (p *PowerPMAC) GetHomePosition(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].HomePos", id)
	return p.RequestFloat(command)
}

// Gets the instantaneous following error
func (p *PowerPMAC) GetPositionError(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].PosError", id)
	return p.RequestFloat(command)
}

// Gets the actual velocity of the motor
func (p *PowerPMAC) GetActualVelocity(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].ActVel", id)
	return p.RequestFloat(command)
}

// Gets the currently commanded desired velocity
func (p *PowerPMAC) GetDesiredVelocity(id int) (float64, error) {
	command := fmt.Sprintf("Motor[%d].DesVel", id)
	return p.RequestFloat(command)
}

// Indicates if the motor's following error is within the allowed band
// and desired velocity is zero
func (p *PowerPMAC) IsInPosition(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].InPos", id)
	return p.RequestBool(command)
}

// Indicates if the commanded trajectory has reached zero velocity
func (p *PowerPMAC) IsDesiredVelocityZero(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].DesVelZero", id)
	return p.RequestBool(command)
}

// Indicates that a homing-search move has succesfully finished
func (p *PowerPMAC) IsHomeComplete(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].HomeComplete", id)
	return p.RequestBool(command)
}

// Indicates that a homing-search is currently active
func (p *PowerPMAC) IsHomeInProgress(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].HomeInProgress", id)
	return p.RequestBool(command)
}

// Gets the status of hardware plus limit switch
func (p *PowerPMAC) IsPlusLimitActive(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].PlusLimit ", id)
	return p.RequestBool(command)
}

// Gets the status of hardware minus limit switch
func (p *PowerPMAC) IsMinusLimitActive(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].MinusLimit ", id)
	return p.RequestBool(command)
}

// Gets the status of software plus limit switch
func (p *PowerPMAC) IsSoftPlusLimitActive(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].SoftPlusLimit", id)
	return p.RequestBool(command)
}

// Gets the status of software minus limit switch
func (p *PowerPMAC) IsSoftMinusLimitActive(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].SoftMinusLimit ", id)
	return p.RequestBool(command)
}
