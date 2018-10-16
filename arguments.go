package main

import "flag"

type cliParams struct {
	log    string
	config string
}

func parseArguments() cliParams {
	var params cliParams
	logFile := flag.String("logfile", "", "Path to the log")
	config := flag.String("config", "", "Path to config file")
	flag.Parse()
	params.log = *logFile
	params.config = *config
	return params
}
