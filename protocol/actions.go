package protocol

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type JogMode string

const (
	JogPlus  JogMode = "+"
	JogMinus JogMode = "-"
	JogStop  JogMode = "/"
)

func motorIDs(ids []int) (string, error) {
	if len(ids) == 0 {
		return "", errors.New("at least one motor id is required")
	}
	parts := make([]string, len(ids))
	for i, id := range ids {
		parts[i] = strconv.Itoa(id)
	}
	return strings.Join(parts, ","), nil
}

// Start/Stop an indefinite jogging
func (p *PowerPMAC) Jog(mode JogMode, ids ...int) error {
	m, err := motorIDs(ids)
	if err != nil {
		return err
	}
	_, err = p.Request(fmt.Sprintf("#%sjog%s", m, mode))
	return err
}

// Starts an absolute jog to a specific position
func (p *PowerPMAC) JogAbsolute(position float64, ids ...int) error {
	m, err := motorIDs(ids)
	if err != nil {
		return err
	}
	_, err = p.Request(fmt.Sprintf("#%sjog=%f", m, position))
	return err
}

// Starts a relative jog by a specified distance from the current
// commanded position
func (p *PowerPMAC) JogRelative(distance float64, ids ...int) error {
	m, err := motorIDs(ids)
	if err != nil {
		return err
	}
	_, err = p.Request(fmt.Sprintf("#%sjog:%f", m, distance))
	return err
}

// Start an absolute jog for the specified coordinate system
func (p *PowerPMAC) MoveAxisAbsolute(coord int, axis string, target float64) error {
	_, err := p.Request(fmt.Sprintf("&%dcx%s%f", coord, axis, target))
	return err
}

// Starts a homing-search move to establish an absolute position reference
func (p *PowerPMAC) Home(ids ...int) error {
	m, err := motorIDs(ids)
	if err != nil {
		return err
	}
	_, err = p.Request(fmt.Sprintf("#%shome", m))
	return err
}

// Establishes the current position as zero without motion
func (p *PowerPMAC) ZeroMoveHome(ids ...int) error {
	m, err := motorIDs(ids)
	if err != nil {
		return err
	}
	_, err = p.Request(fmt.Sprintf("#%shomez", m))
	return err
}

// Immediately disables servo control, opens the loop, and cuts power
func (p *PowerPMAC) Kill(ids ...int) error {
	m, err := motorIDs(ids)
	if err != nil {
		return err
	}
	_, err = p.Request(fmt.Sprintf("#%skill", m))
	return err
}

// Engaging the brake and waiting before disabling power
func (p *PowerPMAC) DelayedKill(ids ...int) error {
	m, err := motorIDs(ids)
	if err != nil {
		return err
	}
	_, err = p.Request(fmt.Sprintf("#%sdkill", m))
	return err
}

// Starts the execution of a Script PLC Program
func (p *PowerPMAC) EnablePLC(name string) error {
	_, err := p.Request(fmt.Sprintf("enable plc %s", name))
	return err
}
