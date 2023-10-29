package outputparser

import (
	"Best-Wifi-Connector/globals"
	"Best-Wifi-Connector/models"
	"Best-Wifi-Connector/utilities"
	"errors"
	"strings"

	"github.com/kr/pretty"
	"go.uber.org/zap"
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

func NMCLIDeviceDetails(input [][]string) []*models.DeviceDetails {
	result := make([]*models.DeviceDetails, 0)
	for _, d := range input {
		obj, err := NMCLIDeviceDetail(d)
		if err != nil {
			globals.Logger.Warn("unable to parse device details", zap.String("err", err.Error()), zap.String("unparsed string", pretty.Sprint(d)))
			continue
		}
		result = append(result, obj)

	}
	return result
}

// Gives parsed output for command : nmcli -t dev status for a single row.
func NMCLIDeviceDetail(input []string) (*models.DeviceDetails, error) {

	if len(input) == 4 {
		return &models.DeviceDetails{
			DeviceID:     input[0],
			DeviceType:   input[1],
			DeviceStatus: (input[2]),
			ConnectedTo:  input[3],
		}, nil

	}

	if len(input) == 3 {
		return &models.DeviceDetails{
			DeviceID:     input[0],
			DeviceType:   input[1],
			DeviceStatus: (input[2]),
		}, nil
	}

	if len(input) == 2 {
		return &models.DeviceDetails{
			DeviceID:   input[0],
			DeviceType: input[1],
		}, nil
	}

	return nil, errors.New("cannot parse input, not enough items")
}
