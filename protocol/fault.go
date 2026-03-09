package protocol

import "fmt"

// Indicates that a fatal following error ocurred
func (p *PowerPMAC) HasFatalFollowingError(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].FeFatal", id)
	return p.RequestBool(command)
}

// Indicates that a warning following error limit is exceeded
func (p *PowerPMAC) HasWarningFollowingError(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].FeWarn", id)
	return p.RequestBool(command)
}

// Indicates that the amplifier has reported a fault signal
func (p *PowerPMAC) HasAmplifierFault(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].AmpFault", id)
	return p.RequestBool(command)
}

// Indicates that the feedback sensor signal as been lost or is invalid
func (p *PowerPMAC) HasEncoderFault(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].EncLoss", id)
	return p.RequestBool(command)
}

// Indicates a secondary fault detection
func (p *PowerPMAC) HasAuxiliarFault(id int) (bool, error) {
	command := fmt.Sprintf("Motor[%d].AuxFault", id)
	return p.RequestBool(command)
}
