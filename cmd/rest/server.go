package main

import "os"

func goMain(args []string) int {

	return 0
}

func main() {
	var status = goMain(os.Args)
	os.Exit(status)
}
