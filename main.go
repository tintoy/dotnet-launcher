package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

// AppName is the name of the application (as launched).
var AppName string

// OutputPrefix is the prefix attached to all program output.
var OutputPrefix string

func main() {
	AppName = path.Base(os.Args[0])
	OutputPrefix = fmt.Sprintf("%s (dotnet-launcher)", AppName)

	log.SetPrefix(OutputPrefix)
	log.SetFlags(0) // No other prefix

	entryAssembly, err := GetEntryAssembly()
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Cannot find entry-point assembly '%s'.\n", entryAssembly)

			os.Exit(254)
		}

		log.Fatal(err)
	}

	// "dotnet MyApp.dll"
	exitCode, err := RunDotNetCLI(entryAssembly)
	if err != nil {
		log.Fatal(err)
	}

	// Pass on the exit code.
	os.Exit(exitCode)
}
