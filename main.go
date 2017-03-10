package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/kardianos/osext"
)

func main() {
	thisExecutable, err := osext.Executable()
	if err != nil {
		log.Fatal(err)
	}

	entryAssembly := strings.TrimSuffix(thisExecutable, ".exe") + ".dll"

	_, err = os.Stat(entryAssembly)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Cannot find entry-point assembly '%s'.", entryAssembly)

			os.Exit(254)
		}

		log.Fatal(err)
	}

	dotnetToolPath, err := exec.LookPath("dotnet")
	if err != nil {
		log.Fatal(err)
	}

	dotnetTool := exec.Command(dotnetToolPath, entryAssembly)
	dotnetTool.Stdin = os.Stdin
	dotnetTool.Stdout = os.Stdout
	dotnetTool.Stderr = os.Stderr

	err = dotnetTool.Run()
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if !ok {
			log.Fatalf("dotnet-launcher: %s",
				err.Error(),
			)
		}
	}

	// Pass on the exit code.
	os.Exit(
		getExitCode(dotnetTool),
	)
}

func getExitCode(command *exec.Cmd) int {
	rawStatus := command.ProcessState.Sys()

	waitStatus, ok := rawStatus.(syscall.WaitStatus)
	if !ok {
		log.Fatalf(
			"dotnet-launcher: Unable to determine process exit code (got unexpected process state '%#v')", rawStatus,
		)
	}

	return waitStatus.ExitStatus()
}
