package motd

import (
	"fmt"
	"runtime"
	"strings"

	"os/exec"
)

var Version string = ""
var Module string = ""
var Runtime string = runtime.Version()

func Info() {
	if len(Module) == 0 {
		cmd := exec.Command("go", "list", "-m")
		if output, err := cmd.Output(); err == nil {
			Module = strings.TrimRight(string(output), "\n")
		}
	}

	if len(Version) == 0 {
		cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
		if output, err := cmd.Output(); err == nil {
			Version = strings.TrimRight(string(output), "\n")
		}
	}

	const info string = "\nModule\t\t%s\nVersion\t\t%s\nRuntime\t\t%s\n\n"

	fmt.Printf(info, Module, Version, Runtime)
}
