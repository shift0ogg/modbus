package modbus

import (
	"testing"
	"time"

	"github.com/shift0ogg/modbus"
)

func Test1(t *testing.T) {
	var c modbus.Client
	if true {
		handler := modbus.NewTCPClientHandler("127.0.0.1:502")
		handler.Timeout = 10 * time.Second
		handler.SlaveId = 0x1
		c = modbus.NewClient(handler)
	} else {
		//
		handler := modbus.NewRTUClientHandler("COM1")
		handler.BaudRate = 9600
		handler.DataBits = 8
		handler.Parity = "N"
		handler.StopBits = 1
		handler.SlaveId = 0x1
		handler.Timeout = 1000 * time.Millisecond
		c = modbus.NewClient(handler)
	}

	bytes, err := c.ReadHoldingRegisters(0, 10)
	if err != nil {
		t.Error(bytes)
	}
	t.Logf("%v", bytes)
	t.Log(bytes)
	c.Close()

}
