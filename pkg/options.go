package gomiko

import (
	"github.com/0x234/gomiko/pkg/types"
)

type DeviceOption func(interface{}) error

func SecretOption(secret string) func(device interface{}) error {
	return func(device interface{}) error {
		device.(types.CiscoDevice).SetSecret(secret)
		return nil
	}
}

func TimeoutOption(timeout uint8) func(device interface{}) error {
	return func(device interface{}) error {
		device.(types.CiscoDevice).SetTimeout(timeout)
		return nil
	}
}
