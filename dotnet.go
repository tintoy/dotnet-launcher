package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// RunDotNetCLI executes the "dotnet" CLI with the specified arguments.
func RunDotNetCLI(args ...string) (exitCode int, err error) {
	var (
		toolPath string
		tool     *exec.Cmd
	)

	toolPath, err = exec.LookPath("dotnet")
	if err != nil {
		return
	}

	tool = exec.Command(toolPath, args...)
	tool.Stdin = os.Stdin
	tool.Stdout = os.Stdout
	tool.Stderr = os.Stderr

	err = tool.Run()
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if !ok {
			return
		}
	}

	exitCode, err = getExitCode(tool)

	return
}

// Retrieves the exit code for the specified command (if available).
func getExitCode(command *exec.Cmd) (exitCode int, err error) {
	if !command.ProcessState.Exited() {
		err = fmt.Errorf(
			"Process %d has not exited yet", command.ProcessState.Pid(),
		)

		return
	}
	rawStatus := command.ProcessState.Sys()

	waitStatus, ok := rawStatus.(syscall.WaitStatus)
	if !ok {
		err = fmt.Errorf(
			"Unable to determine process exit code (got unexpected process state '%#v')", rawStatus,
		)

		return
	}

	exitCode = waitStatus.ExitStatus()

	return
}
