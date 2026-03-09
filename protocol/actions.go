package protocol

import "fmt"

type JodeMode string

const (
	JogPlus  = "+"
	JogMinus = "-"
	JogStop  = "/"
)

// Start/Stop an indefinite jogging
func (p *PowerPMAC) Jog(id int, mode JodeMode) error {
	command := fmt.Sprintf("#%djog%s", id, mode)
	_, err := p.Request(command)
	return err
}

// Starts an absolute jog to a specific position
func (p *PowerPMAC) JogAbsolute(id int, position float64) error {
	command := fmt.Sprintf("#%djog=%f", id, position)
	_, err := p.Request(command)
	return err
}

// Starts a relative jog by a specified distance from the current
// commanded position
func (p *PowerPMAC) JogRelative(id int, distance float64) error {
	command := fmt.Sprintf("#%djog:%f", id, distance)
	_, err := p.Request(command)
	return err
}

// Start an absolute jog for the specified coordinate system
func (p *PowerPMAC) MoveAxisAbsolute(coord int, axis string, target float64) error {
	command := fmt.Sprintf("&%dcx%s%f", coord, axis, target)
	_, err := p.Request(command)
	return err
}

// Starts a homing-search move to establish an absolute position reference
func (p *PowerPMAC) Home(id int) error {
	command := fmt.Sprintf("#%dhome", id)
	_, err := p.Request(command)
	return err
}

// Establishes the current position as zero without motion
func (p *PowerPMAC) ZeroMoveHome(id int) error {
	command := fmt.Sprintf("#%dhomez", id)
	_, err := p.Request(command)
	return err
}

// Immediately disables servo control, opens the loop, and cuts power
func (p *PowerPMAC) Kill(id int) error {
	command := fmt.Sprintf("#%dkill", id)
	_, err := p.Request(command)
	return err
}

// Engaging the brake and waiting before disabling power
func (p *PowerPMAC) DelayedKill(id int) error {
	command := fmt.Sprintf("#%ddkill", id)
	_, err := p.Request(command)
	return err
}

// Starts the execution of a Script PLC Program
func (p *PowerPMAC) EnablePLC(name string) error {
	command := fmt.Sprintf("enable plc %s", name)
	_, err := p.Request(command)
	return err
}
