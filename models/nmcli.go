package models

type DeviceDetails struct {
	DeviceID     string
	DeviceType   string
	DeviceStatus string
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
