package interfaces

type WifiControllerInterface interface {
	GetActiveConnections() bool // model to be changed
	GetSavedNetworks() bool     // Model to be changed
	ConnectToNetwork(ssid string) bool
	DisconnectFromNetwork(ssid string) bool
	IsWifiWorking() bool
	EnableWifiModule() bool
	DisableWifiModule() bool
}
