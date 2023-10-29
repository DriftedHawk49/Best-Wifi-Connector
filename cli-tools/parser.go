package clitools

import "flag"

type Flags struct {
	Debug bool
}

func FlagParser() *Flags {

	d := flag.Bool("d", false, "enable debug mode, which prints debug logs while running the tool. Useful in development mode")

	flag.Parse()

	return &Flags{
		Debug: *d,
	}

}
