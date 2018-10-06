package main

import "fmt"
import "flag"

func main() {
	/* This is my first sample program. */
	logFile := flag.String("logfile", "", "Path to the log")
	flag.Parse()
	fmt.Printf("The log is : %s\n", *logFile)
}
