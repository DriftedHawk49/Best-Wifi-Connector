package parser

const (
	DEVICE_STATUS_CONNECTED    deviceStatus = "connected"
	DEVICE_STATUS_DISCONNECTED deviceStatus = "disconnected"
	DEVICE_STATUS_UNMANAGED    deviceStatus = "unmanaged"
)

type deviceStatus string

type DeviceDetails struct {
	DeviceID     string
	DeviceType   string
	DeviceStatus deviceStatus
	ConnectedTo  string
}
