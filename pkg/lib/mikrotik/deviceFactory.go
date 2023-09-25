package mikrotik

import (
	"errors"
	"github.com/0x234/gomiko/pkg/connections"
	"github.com/0x234/gomiko/pkg/driver"
	"github.com/0x234/gomiko/pkg/types"
)

func NewDevice(connection connections.Connection, DeviceType string) (types.Device, error) {
	devDriver := driver.NewDriver(connection, "\r")

	switch DeviceType {
	case "mikrotik_routeros":
		return &MikroTikRouterOS{
			Driver:     devDriver,
			DeviceType: DeviceType,
			Prompt:     "",
		}, nil
	default:
		return nil, errors.New("unsupported DeviceType: " + DeviceType)

	}

}
