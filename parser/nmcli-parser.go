package parser

import (
	"Best-Wifi-Connector/utilities"
	"errors"
	"os/exec"
	"strings"
)

// Expects tersed output, first breaks down by line, then breaks down by : (colon)
func NMCLIParser(output []byte) [][]string {

	result := make([][]string, 0)

	outputString := string(output)
	outputByLine := strings.Split(outputString, "\n")

	for _, l := range outputByLine {
		o := utilities.SplitEscapeSafeString(l, ":")
		result = append(result, o)
	}

	return result

}

// Gives parsed output for command : nmcli -t dev status for a single row.
func NMCLIDeviceDetails(input []string) (*DeviceDetails, error) {

	if len(input) == 4 {
		return &DeviceDetails{
			DeviceID:     input[0],
			DeviceType:   input[1],
			DeviceStatus: deviceStatus(input[2]),
			ConnectedTo:  input[3],
		}, nil

	}

	if len(input) == 3 {
		return &DeviceDetails{
			DeviceID:     input[0],
			DeviceType:   input[1],
			DeviceStatus: deviceStatus(input[2]),
		}, nil
	}

	if len(input) == 2 {
		return &DeviceDetails{
			DeviceID:   input[0],
			DeviceType: input[1],
		}, nil
	}

	return nil, errors.New("cannot parse input, not enough items")
}

// Tells whether system wifi is enabled or not
func NMCLIIsWifiConnected() (bool, error) {
	output, err := exec.Command("nmcli", "-t", "radio", "wifi").Output()

	if err != nil {
		return false, err
	}

	return string(output) == string(WIFI_ENABLED), nil
}

// Turns on the wifi if it is disabled
func NMCLITurnOnWifi() error {
	_, err := exec.Command("nmcli", "-t", "radio", "wifi", "on").Output()
	return err
}
