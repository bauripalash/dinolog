package main

import (
	"dinolog/parser"
	"log"
	"os"
)

var logfile = "example.dinolog.txt"

func main() {
	if _, err := os.Stat(logfile); err != nil {
		log.Fatalf("Dinologg file %s does not exist", logfile)
		//fmt.Println(
	}
	data, err := os.ReadFile(logfile)

	if err != nil {
		log.Fatalln("Failed to read dinolog file ->" + logfile)
	}

	parser.Parse(string(data))
}
