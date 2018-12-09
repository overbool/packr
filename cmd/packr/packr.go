package main

import "os"

func main() {
	rootCMD.AddCommand(versionCMD)
	if err := rootCMD.Execute(); err != nil {
		os.Exit(1)
	}
}
