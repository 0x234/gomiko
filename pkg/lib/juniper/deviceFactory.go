package juniper

import (
	"github.com/0x234/gomiko/pkg/connections"
	"github.com/0x234/gomiko/pkg/driver"
	"github.com/0x234/gomiko/pkg/types"
)

func NewDevice(connection connections.Connection, DeviceType string) (types.Device, error) {
	devDriver := driver.NewDriver(connection, "\n")
	return &JunOSDevice{
		Prompt: "",
		Driver: devDriver,
	}, nil

}
