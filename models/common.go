package models

type ActiveConnections struct {
	Mode           string // 2.4GHZ or 5GHz , use ENUM
	Connected      bool   // Whether system is connected to this network?
	SSID           string // Human readable name of the network
	BSSID          string // Network identifier string
	Channel        int    // Channel to which it is connected
	SignalStrength int    // 0-100 Signal Strength
}

type SavedConnections struct {
	Mode      string // 2.4GHz or 5GHz, use ENUM
	Connected bool   // Whether system is connected to this network?
	SSID      string // Human readable name of the network
	BSSID     string // Network identifier string
}
