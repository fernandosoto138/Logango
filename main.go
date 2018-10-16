package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkFile(err error, file string) {
	if err != nil {
		fmt.Println("Problems loading the file " + file)
		panic(err)

	}
}

func main() {
	args := parseArguments()
	config := NewConfig(args.config)
	if config == nil {
		panic("Problems loading the config file " + args.config)
	}
	fmt.Printf("%v", config)
	file, err := os.Open(args.log)
	checkFile(err, args.log)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
