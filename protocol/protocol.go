/*
Author: Leonardo Rossi Leao
Created at: March 06th, 2025
*/

package protocol

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	CR  byte = 0x0D
	LF  byte = 0x0A
	ACK byte = 0x06
)

type SSHOptions struct {
	client  *ssh.Client
	session *ssh.Session
	stdin   io.WriteCloser
	stdout  io.Reader

	Host     string
	Port     int
	Username string
	Password string
	Timeout  time.Duration
}

type PowerPMAC struct {
	mutex sync.Mutex
	SSH   SSHOptions
}

// Establishes a connection with Power PMAC through SSH
func (p *PowerPMAC) Connect() error {
	var err error

	p.mutex.Lock()
	defer p.mutex.Unlock()

	defer func() {
		if err != nil {
			p.closeUnsafe()
		}
	}()

	url := fmt.Sprintf("%s:%d", p.SSH.Host, p.SSH.Port)
	p.SSH.client, err = ssh.Dial("tcp", url, &ssh.ClientConfig{
		User: p.SSH.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(p.SSH.Password),
		},
		Timeout:         p.SSH.Timeout,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		return err
	}

	p.SSH.session, err = p.SSH.client.NewSession()
	if err != nil {
		return err
	}
	if err = p.SSH.session.RequestPty("vt100", 80, 40, ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}); err != nil {
		return err
	}
	if p.SSH.stdin, err = p.SSH.session.StdinPipe(); err != nil {
		return err
	}
	if p.SSH.stdout, err = p.SSH.session.StdoutPipe(); err != nil {
		return err
	}
	if err = p.SSH.session.Shell(); err != nil {
		return err
	}

	output := make([]byte, 1024)
	_, err = p.readUnsafe(output, []byte("ppmac#"))
	if _, err = p.writeUnsafe("gpascii -2"); err != nil {
		return err
	}

	output = make([]byte, 1024)
	_, err = p.readUnsafe(output, []byte("STDIN Open for ASCII Input\r\n"))
	return err

}

func (p *PowerPMAC) closeUnsafe() error {
	var errs []error

	if p.SSH.stdout != nil {
		p.SSH.stdout = nil
	}
	if p.SSH.stdin != nil {
		errs = append(errs, p.SSH.stdin.Close())
		p.SSH.stdin = nil
	}
	if p.SSH.session != nil {
		errs = append(errs, p.SSH.session.Close())
		p.SSH.session = nil
	}
	if p.SSH.client != nil {
		errs = append(errs, p.SSH.client.Close())
		p.SSH.client = nil
	}

	return errors.Join(errs...)
}

// Closes the connection with Power PMAC
func (p *PowerPMAC) Close() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.closeUnsafe()
}

func (p *PowerPMAC) isConnectedUnsafe() bool {
	if p.SSH.client == nil {
		return false
	}
	_, _, err := p.SSH.client.SendRequest("keepalive@openssh.com", true, nil)
	return err == nil
}

// Checks if a connection is established
func (p *PowerPMAC) IsConnected() bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.isConnectedUnsafe()
}

func (p *PowerPMAC) writeUnsafe(command string) (int, error) {
	if !p.isConnectedUnsafe() {
		return 0, fmt.Errorf("Power PMAC is not connected")
	}
	commandBytes := []byte(command)
	if commandBytes[len(commandBytes)-1] != CR {
		commandBytes = append(commandBytes, CR)
	}
	return p.SSH.stdin.Write(commandBytes)
}

// Writes a command to power PMAC
func (p *PowerPMAC) Write(command string) (int, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.writeUnsafe(command)
}

func (p *PowerPMAC) readUnsafe(output []byte, until ...[]byte) (int, error) {
	if !p.isConnectedUnsafe() {
		return 0, fmt.Errorf("Power PMAC is not connected")
	}

	terminators := [][]byte{{CR, LF, ACK}}
	if len(until) > 0 {
		terminators = until
	}

	totalRead := 0
	deadline := time.Now().Add(p.SSH.Timeout)
	for {
		if time.Now().After(deadline) {
			return totalRead, fmt.Errorf("power pmac read timeout")
		}
		numBytes, err := p.SSH.stdout.Read(output[totalRead:])
		if numBytes > 0 {
			totalRead += numBytes
			for _, terminator := range terminators {
				if bytes.Contains(output[:totalRead], terminator) {
					if bytes.Contains(output[:totalRead], []byte("error #")) {
						return totalRead, fmt.Errorf("power pmac error: %s", string(output[:totalRead]))
					}
					return totalRead, nil
				}
			}
			if totalRead >= len(output) {
				return totalRead, fmt.Errorf("output buffer full (%d bytes)", len(output))
			}
		}
		if err != nil {
			if err == io.EOF {
				return totalRead, nil
			}
			return totalRead, fmt.Errorf("ssh read error: %w", err)
		}
		time.Sleep(1 * time.Millisecond)
	}
}

// Reads a message from Power PMAC
func (p *PowerPMAC) Read(output []byte) (int, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.readUnsafe(output)
}

// Requests a command from Power PMAC
func (p *PowerPMAC) Request(command string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if _, err := p.writeUnsafe(command); err != nil {
		return "", err
	}
	output := make([]byte, 1024)
	n, err := p.readUnsafe(output)
	if err != nil {
		return "", err
	}
	return string(output[:n]), nil
}

// Requests a float value from Power PMAC
func (p *PowerPMAC) RequestFloat(command string) (float64, error) {
	response, err := p.Request(command)
	if err != nil {
		return 0, err
	}
	trimmed := strings.TrimRight(response, "\r\n\x06")
	strValue := trimmed
	if idx := strings.Index(trimmed, "="); idx != -1 {
		strValue = trimmed[idx+1:]
	}
	return strconv.ParseFloat(strValue, 64)
}

// Requests an int value from Power PMAC
func (p *PowerPMAC) RequestInt(command string) (int, error) {
	response, err := p.Request(command)
	if err != nil {
		return 0, err
	}
	trimmed := strings.TrimRight(response, "\r\n\x06")
	strValue := trimmed
	if idx := strings.Index(trimmed, "="); idx != -1 {
		strValue = trimmed[idx+1:]
	}
	return strconv.Atoi(strValue)
}

// Request a boolean value from Power PMAC
func (p *PowerPMAC) RequestBool(command string) (bool, error) {
	value, err := p.RequestInt(command)
	if err != nil {
		return false, err
	}
	return value == 1, nil
}
