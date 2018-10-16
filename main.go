package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

func checkFile(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	args := parseArguments()
	file, err := os.Open(args.log)
	checkFile(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Printf(scanner.Text())
		fmt.Printf("\n")
	}
}
