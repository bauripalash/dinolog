package parser

import (
	"fmt"
	"strings"
)

type MetaData struct {
	Title  string
	Author string
	Others map[string]string
}

type BlogItem struct {
	Date     string
	Text     string
	BlogMeta map[string]string
}

func readOption(line string) (string, string, bool) {
	temp := strings.Split(line, "=")
	if len(temp) >= 2 {
		return strings.TrimSpace(temp[0]), strings.TrimSpace(strings.Join(temp[1:], "=")), true
	}

	return "", "", false
}

func popLine(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func Parse(src string) {
	data := strings.Split(src, "\n")
	md := MetaData{Others: make(map[string]string)}

	for index, item := range data {
		if strings.HasPrefix(item, "----") {
			break
		}

		fmt.Printf("%d |%s\n", index, item)
		if key, value, ok := readOption(item); ok {

			switch key {
			case "title":
				md.Title = value
			case "author":
				md.Author = value
			default:
				md.Others[key] = value
			}

		}

	}

	fmt.Printf("%#v\n\n", md)
	fmt.Println(strings.Join(data, "\n"))
}
