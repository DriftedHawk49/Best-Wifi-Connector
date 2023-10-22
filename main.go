package main

import (
	"Best-Wifi-Connector/parser"
	"Best-Wifi-Connector/utilities"
	"fmt"
	"log"
	"os/exec"

	"github.com/kr/pretty"
)

func main() {

	output, err := exec.Command("nmcli", "-t", "dev", "wifi", "list").Output()

	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(pretty.Sprint(utilities.TrimEscapeCharacters(parser.NMCLIParser(output)[0][1])))

}
