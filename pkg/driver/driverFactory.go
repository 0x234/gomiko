package driver

import "github.com/0x234/gomiko/pkg/connections"

type IDriver interface {
	Connect() error
	Disconnect()
	SendCommand(cmd string, expectPattern string) (string, error)
	SendCommandsSet(cmds []string, expectPattern string) (string, error)
	FindDevicePrompt(regex string, pattern string) (string, error)
	ReadUntil(pattern string) (string, error)
	SetTimeout(timeout uint8)
}

func NewDriver(Connection connections.Connection, Return string) IDriver {
	return &Driver{
		Connection: Connection,
		Return:     Return,
	}

}
