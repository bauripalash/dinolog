package parser

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MetaData struct {
	Title  string
	Author string
	Others map[string]string
}

type BlogData struct {
	Meta    MetaData
	Follows map[string]string
	Posts   []BlogItem
}

type BlogItem struct {
	Date     string
	Text     string
	BlogMeta map[string]string
}

func pprint(b *BlogData) {
	ijson, err := json.MarshalIndent(b, "", "  ")

	if err != nil {
		return
	}

	fmt.Printf("%s\n", ijson)
}

func readOption(line string) (string, string, bool) {
	temp := strings.Split(line, "=")
	if len(temp) >= 2 {
		return strings.TrimSpace(temp[0]), strings.TrimSpace(strings.Join(temp[1:], "=")), true
	}

	return "", "", false
}

func popLines(slice []string, index int) []string {
	return slice[index:]
}

func Parse(src string) {
	data := strings.Split(src, "\n")
	blg := BlogData{Meta: MetaData{Others: make(map[string]string)}, Follows: make(map[string]string), Posts: make([]BlogItem, 0)}
	lastindex := 0
	//params parsing
	for _, item := range data {
		if strings.HasPrefix(item, "##") {
			continue
		}
		if strings.HasPrefix(item, "----") {
			break
		}

		//fmt.Printf("%d |%s\n", index, item)
		if key, value, ok := readOption(item); ok {

			switch key {
			case "title":
				blg.Meta.Title = value
			case "author":
				blg.Meta.Author = value
			default:
				blg.Meta.Others[key] = value
			}
			lastindex += 1
		}

	}
	data = popLines(data, lastindex)
	isInsideFollowList := false
	for index, item := range data {
		if strings.HasPrefix(item, "----") {
			isInsideFollowList = !isInsideFollowList
			if !isInsideFollowList {
				lastindex = index + 1
				break
			}
		}

		if isInsideFollowList {
			if k, v, ok := readOption(item); ok {
				if strings.HasPrefix(k, "@") {
					blg.Follows[k] = v
				}
			}
		}

	}

	data = popLines(data, lastindex)
	posts := []BlogItem{}

	tempItem := BlogItem{BlogMeta: make(map[string]string)}
	insideBlogMeta := false
	blogMetaLock := false
	for _, item := range data {
		item = strings.TrimSpace(item)
		if strings.HasPrefix(item, "[[") { //Read Dates
			if strings.HasSuffix(item, "]]") {
				dt := strings.TrimPrefix(item, "[[")
				dt = strings.TrimSuffix(dt, "]]")
				tempItem.Date = dt
				continue
			}
		}

		if !strings.HasPrefix(item, "--0--") { // This Means post has not yet ended

			if strings.HasPrefix(item, "++++") && !blogMetaLock { //Check for post metas
				insideBlogMeta = !insideBlogMeta
				if !insideBlogMeta {
					blogMetaLock = true //One post can have only one meta block
				}
				continue
			}

			if insideBlogMeta && !blogMetaLock {
				if k, v, ok := readOption(item); ok {
					tempItem.BlogMeta[k] = v
				}
			} else {
				tempItem.Text += item + "\n"
			}

		} else if strings.HasPrefix(item, "--0--") {
			posts = append(posts, tempItem)
			tempItem = BlogItem{BlogMeta: make(map[string]string)}
			blogMetaLock = false
			continue
		}

	}

	blg.Posts = posts

	//fmt.Printf("%#v\n", posts)

	//fmt.Printf("%#v\n\n", md)
	pprint(&blg)
	//fmt.Println(strings.Join(data, "\n"))
}
