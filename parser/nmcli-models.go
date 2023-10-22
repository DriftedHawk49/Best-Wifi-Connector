package parser

const (
	DEVICE_STATUS_CONNECTED    deviceStatus = "connected"
	DEVICE_STATUS_DISCONNECTED deviceStatus = "disconnected"
	DEVICE_STATUS_UNMANAGED    deviceStatus = "unmanaged"
	WIFI_ENABLED               wifiStatus   = "enabled"
	WIFI_DISABLED              wifiStatus   = "disabled"
)

type deviceStatus string
type wifiStatus string

type DeviceDetails struct {
	DeviceID     string
	DeviceType   string
	DeviceStatus deviceStatus
	ConnectedTo  string
}

type NetworkDetails struct {
	Connected      bool
	BSSID          string
	SSID           string
	Mode           string
	CHAN           int
	Rate           int
	SignalStrength int
}

type SavedNetwork struct {
	SSID     string
	UUID     string
	Type     string
	DeviceId string
}
