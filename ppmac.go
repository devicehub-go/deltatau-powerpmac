package ppmac

import (
	"time"

	"github.com/devicehub-go/deltatau-powerpmac/protocol"
)

type PowerPMAC = protocol.PowerPMAC
type Options struct {
	Host     string
	Port     int
	Username string
	Password string
	Timeout  time.Duration
}

/*
Creates a new instance of Omron/Delta Tau Power PMAC
allows to communicate and control the connected motors
*/
func New(options Options) *PowerPMAC {
	powerpmac := &PowerPMAC{
		SSH: protocol.SSHOptions{
			Host:     options.Host,
			Port:     options.Port,
			Username: options.Username,
			Password: options.Password,
			Timeout:  options.Timeout,
		},
	}
	return powerpmac
}
