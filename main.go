package main

import (
	"flag"
)

type cliParams struct {
	log string
}

func parseArguments() cliParams {
	var params cliParams
	logFile := flag.String("logfile", "", "Path to the log")
	flag.Parse()
	params.log = *logFile
	return params
}

func main() {
	args := parseArguments()
	println(args.log)
}
