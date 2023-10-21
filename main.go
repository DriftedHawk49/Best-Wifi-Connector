package main

import (
	"Best-Wifi-Connector/parser"
	"fmt"
	"log"
	"os/exec"

	"github.com/kr/pretty"
)

func main() {

	output, err := exec.Command("nmcli", "-t", "dev", "status").Output()

	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(pretty.Sprint(parser.NMCLIDeviceDetails(parser.NMCLIParser(output)[0])))

}
