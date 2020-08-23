package process

import (
	"github.com/mitchellh/go-ps"
)

//Process properties
type Process struct {
	Pid  int    `json:"pid"`
	Name string `json:"name"`
}

//Info returns []Process containing system processes information
func Info() []Process {
	processes, err := ps.Processes()
	if err != nil {
		panic(err)
	}
	var processList []Process
	for _, p := range processes {
		proc := Process{Pid: p.Pid(), Name: p.Executable()}
		processList = append(processList, proc)
	}

	return processList
}
