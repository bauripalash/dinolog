package parser

import (
	"os"
)

func GetDemoObj(logfile string) *BlogData {
	if _, err := os.Stat(logfile); err != nil {
		//log.Fatalf("Dinologg file %s does not exist", logfile)
		return nil
		//fmt.Println(
	}
	data, err := os.ReadFile(logfile)

	if err != nil {
		//log.Fatalln("Failed to read dinolog file ->" + logfile)
		return nil
	}

	return Parse(string(data))

}
