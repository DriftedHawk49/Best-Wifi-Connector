package parser

import (
	"strings"
)

// Expects tersed output, first breaks down by line, then breaks down by : (colon)
func NMCLIParser(output []byte) [][]string {

	result := make([][]string, 0)

	outputString := string(output)
	outputByLine := strings.Split(outputString, "\n")

	for _, l := range outputByLine {
		o := strings.Split(l, ":")
		result = append(result, o)
	}

	return result

}

// Gives parsed output for command : nmcli -t dev status for a single row.
func NMCLIDeviceDetails(input []string) *DeviceDetails {
	result := DeviceDetails{
		DeviceID:     input[0],
		DeviceType:   input[1],
		DeviceStatus: deviceStatus(input[2]),
		ConnectedTo:  input[3],
	}

	return &result
}
