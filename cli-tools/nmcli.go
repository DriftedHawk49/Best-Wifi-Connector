package clitools

import (
	"Best-Wifi-Connector/globals"
	"Best-Wifi-Connector/interfaces"
	"fmt"
	"os/exec"

	"go.uber.org/zap"
)

func NewNMCLI() interfaces.WifiControllerInterface {
	return &NMCLI{
		logger: globals.Logger,
	}
}

type NMCLI struct {
	logger *zap.Logger
}

func (n *NMCLI) IsWifiWorking() bool {
	output, err := exec.Command("nmcli", "radio", "wifi").Output()

	if err != nil {
		n.logger.Error("error while getting wifi status", zap.String("err", err.Error()))
		return false
	}

	return string(output) == globals.WIFI_ENABLED
}

func (n *NMCLI) EnableWifiModule() bool {
	output, err := exec.Command("nmcli", "radio", "wifi", "on").Output()

	if err != nil {
		n.logger.Error("error while enabling Wifi Module", zap.String("err", err.Error()))
		return false
	}

	n.logger.Debug("output returned", zap.String("output", string(output)))
	return true
}

func (n *NMCLI) DisableWifiModule() bool {
	output, err := exec.Command("nmcli", "radio", "wifi", "off").Output()

	if err != nil {
		n.logger.Error("error while disabling Wifi Module", zap.String("err", err.Error()))
		return false
	}

	n.logger.Debug("output returned", zap.String("output", string(output)))
	return true
}

// TODO: Complete this function
func (n *NMCLI) GetSavedNetworks() bool {
	output, err := exec.Command("nmcli", "-t", "con", "show").Output()

	if err != nil {
		n.logger.Error("error while disabling Wifi Module", zap.String("err", err.Error()))
		return false
	}

	n.logger.Debug("output returned", zap.String("output", string(output)))
	return true
}

// TODO: Complete this function
func (n *NMCLI) GetActiveConnections() bool {
	output, err := exec.Command("nmcli", "-t", "dev", "wifi", "list").Output()

	if err != nil {
		n.logger.Error("error while disabling Wifi Module", zap.String("err", err.Error()))
		return false
	}

	n.logger.Debug("output returned", zap.String("output", string(output)))
	return true
}

func (n *NMCLI) ConnectToNetwork(ssid string) bool {
	output, err := exec.Command("nmcli", "con", "up", fmt.Sprintf("\"%s\"", ssid)).Output()
	if err != nil {
		n.logger.Error("error while connecting to another saved Wifi Network", zap.String("err", err.Error()), zap.String("ssid", ssid))
		return false
	}

	n.logger.Debug("output returned", zap.String("output", string(output)))
	return true

}

func (n *NMCLI) DisconnectFromNetwork(ssid string) bool {
	output, err := exec.Command("nmcli", "con", "down", fmt.Sprintf("\"%s\"", ssid)).Output()
	if err != nil {
		n.logger.Error("error while disconnecting from a Wifi Network", zap.String("err", err.Error()), zap.String("ssid", ssid))
		return false
	}

	n.logger.Debug("output returned", zap.String("output", string(output)))

	return true
}
