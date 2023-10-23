package main

import (
	"Best-Wifi-Connector/globals"
	"Best-Wifi-Connector/log"
	"Best-Wifi-Connector/parser"
	"Best-Wifi-Connector/utilities"
	"fmt"
	"os/exec"

	"github.com/kr/pretty"
	"go.uber.org/zap"
)

func main() {

	logger := log.InitLogger("")
	globals.Logger = logger

	output, err := exec.Command("nmcli", "-t", "dev", "wifi", "list").Output()

	if err != nil {
		logger.Error("error while getting wifi list", zap.String("err", err.Error()))
		return
	}

	fmt.Println(pretty.Sprint(utilities.TrimEscapeCharacters(parser.NMCLIParser(output)[0][1])))

}
