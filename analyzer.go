package main

import (
	"fmt"
	"regexp"
)

func analizeLine(line string, config *Config) {
	for i, counter := range config.Counters {
		matched, err := regexp.MatchString(counter.Regex, line)
		if err != nil {
			fmt.Printf("%v", err)
		}
		if matched && err == nil {
			config.Counters[i].Count++
		}
	}
}

func printResults(config *Config) {
	fmt.Println("-------- Counters ----------")
	for _, counter := range config.Counters {
		fmt.Printf("%s : %d\n", counter.DisplayName, counter.Count)
	}
}
