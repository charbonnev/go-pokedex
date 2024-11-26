package main

import "os"

func ExitCommand() error {
	os.Exit(0)
	return nil
}
