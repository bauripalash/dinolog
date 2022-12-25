package main

import (
	"dinolog/frontend"
	"dinolog/parser"
	"log"
	"os"
)

var logfile = "example.dinolog.txt"

func main() {
	runserver := true
	runparser := false
	if runparser {
		if _, err := os.Stat(logfile); err != nil {
			log.Fatalf("Dinologg file %s does not exist", logfile)
			//fmt.Println(
		}
		data, err := os.ReadFile(logfile)

		if err != nil {
			log.Fatalln("Failed to read dinolog file ->" + logfile)
		}

		parser.Parse(string(data))
	} else if runserver {
		frontend.Server()
	}

}
