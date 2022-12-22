package main

import (
	"dinolog/parser"
	"log"
	"os"
)

var logfile = "example.dinolog.txt"

func main() {
	data, err := os.ReadFile(logfile)

	if err != nil {
		log.Fatalln("Failed to read dinolog file ->" + logfile)
	}

	parser.Parse(string(data))
}
