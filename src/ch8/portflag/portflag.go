package portflag

import (
	"flag"
	"fmt"
)

// PortFlag struct
type PortFlag struct{ Port }

// Port type
type Port uint

// PortCommandLine commandline
func PortCommandLine(name string, value Port, usage string) *Port {
	f := PortFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Port
}

// Set func
func (p *PortFlag) Set(s string) error {
	var unit string
	var value uint
	fmt.Sscanf(s, "%d%s", &value, &unit) // no error check needed
	p.Port = Port(value)
	return nil
}

func (p Port) String() string { return fmt.Sprintf("%d", p) }
