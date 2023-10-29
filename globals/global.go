package globals

import "go.uber.org/zap"

// Needs to be initialized as soon as logger is created for use in every module
var Logger *zap.Logger

const (
	LOGLEVEL_DEBUG             = "DEBUG"
	LOGLEVEL_INFO              = "INFO"
	MODE_5GHZ                  = "5GHz"
	MODE_2_4GHZ                = "2.4GHz"
	DEVICE_STATUS_CONNECTED    = "connected"
	DEVICE_STATUS_DISCONNECTED = "disconnected"
	DEVICE_STATUS_UNMANAGED    = "unmanaged"
	WIFI_ENABLED               = "enabled"
	WIFI_DISABLED              = "disabled"
)
