package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kardianos/osext"
)

// GetEntryAssembly retrieves the full path to the assembly containing the program entry-point.
//
// The entry-point assembly should be a file with the same name and location as the current program executable, but ending in ".dll" (regardless of platform).
func GetEntryAssembly() (assemblyPath string, err error) {
	var thisExecutable string

	thisExecutable, err = osext.Executable()
	if err != nil {
		return
	}

	// Ensure this isn't a symbolic link.
	thisExecutable, err = filepath.EvalSymlinks(thisExecutable)
	if err != nil {
		return
	}

	assemblyPath = strings.TrimSuffix(thisExecutable, ".exe") + ".dll"
	_, err = os.Stat(assemblyPath)

	return
}
