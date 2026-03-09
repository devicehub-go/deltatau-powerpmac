package ppmac_test

import (
	"fmt"
	"testing"
	"time"

	ppmac "github.com/devicehub-go/deltatau-powerpmac"
)

func TestPowerPMAC_Getters(t *testing.T) {
	drive := ppmac.New(ppmac.Options{
		Host:     "",
		Port:     0,
		Username: "",
		Password: "",
		Timeout:  10 * time.Second,
	})

	if err := drive.Connect(); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer drive.Close()

	motorID := 1

	// Fault & Status Indicators (Booleans)
	t.Run("StatusBits", func(t *testing.T) {
		statusChecks := []struct {
			name string
			fn   func(int) (bool, error)
		}{
			{"FatalFollowingError", drive.HasFatalFollowingError},
			{"WarningFollowingError", drive.HasWarningFollowingError},
			{"AmplifierFault", drive.HasAmplifierFault},
			{"EncoderFault", drive.HasEncoderFault},
			{"AuxiliarFault", drive.HasAuxiliarFault},
			{"InPosition", drive.IsInPosition},
			{"DesiredVelocityZero", drive.IsDesiredVelocityZero},
			{"HomeComplete", drive.IsHomeComplete},
			{"HomeInProgress", drive.IsHomeInProgress},
			{"PlusLimit", drive.IsPlusLimitActive},
			{"MinusLimit", drive.IsMinusLimitActive},
			{"SoftPlusLimit", drive.IsSoftPlusLimitActive},
			{"SoftMinusLimit", drive.IsSoftMinusLimitActive},
		}

		for _, sc := range statusChecks {
			val, err := sc.fn(motorID)
			if err != nil {
				t.Errorf("%s failed: %v", sc.name, err)
			} else {
				fmt.Printf("%s: %v\n", sc.name, val)
			}
		}
	})

	// Real-time Metrics (Floats)
	t.Run("Metrics", func(t *testing.T) {
		metrics := []struct {
			name string
			fn   func(int) (float64, error)
		}{
			{"ActualPosition", drive.GetActualPosition},
			{"DesiredPosition", drive.GetDesiredPosition},
			{"PositionError", drive.GetPositionError},
			{"ActualVelocity", drive.GetActualVelocity},
			{"DesiredVelocity", drive.GetDesiredVelocity},
		}

		for _, m := range metrics {
			val, err := m.fn(motorID)
			if err != nil {
				t.Errorf("%s failed: %v", m.name, err)
			} else {
				fmt.Printf("%s: %f\n", m.name, val)
			}
		}
	})

	// Configuration Settings
	t.Run("Settings", func(t *testing.T) {
		if val, err := drive.GetJogAcceleration(motorID); err == nil {
			fmt.Printf("Jog Acceleration: %f\n", val)
		}
		if val, err := drive.GetJogJerk(motorID); err == nil {
			fmt.Printf("Jog Jerk: %f\n", val)
		}
		if val, err := drive.GetHomeVelocity(motorID); err == nil {
			fmt.Printf("Home Velocity: %f\n", val)
		}
		if val, err := drive.GetHomeOffset(motorID); err == nil {
			fmt.Printf("Home Offset: %f\n", val)
		}
		if val, err := drive.GetMaximumSpeed(motorID); err == nil {
			fmt.Printf("Max Speed: %f\n", val)
		}
		if val, err := drive.GetMaximumPosition(motorID); err == nil {
			fmt.Printf("Max Position: %f\n", val)
		}
		if val, err := drive.GetMininumPosition(motorID); err == nil {
			fmt.Printf("Min Position: %f\n", val)
		}
		if val, err := drive.GetMaximumOutput(motorID); err == nil {
			fmt.Printf("Max DAC Output: %f\n", val)
		}
		if val, err := drive.GetFatalFollowingErrorLimit(motorID); err == nil {
			fmt.Printf("Fatal FE Limit: %f\n", val)
		}
		if val, err := drive.GetWarningFollowingErrorLimit(motorID); err == nil {
			fmt.Printf("Warning FE Limit: %f\n", val)
		}
		if val, err := drive.GetServoControl(motorID); err == nil {
			fmt.Printf("Servo Control Mode: %v\n", val)
		}
		if val, err := drive.GetLeaderMotor(motorID); err == nil {
			fmt.Printf("Leader Motor ID: %d\n", val)
		}
		if val, err := drive.GetInPositionBand(motorID); err == nil {
			fmt.Printf("In Position Band: %f\n", val)
		}
	})
}
