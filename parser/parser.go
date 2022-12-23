package parser

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	DATE_START = "[["
	DATA_END = "]]"
	PAGINATION_NEXT = ">>>"
	PAGINATION_PREV = "<<<"
	POST_END = "--0--"
	COMMENT = "##"
	FOLLOW = "@"
	POST_META = "++++"
	FOLLOW_BLOCK = "----"
	CONFIG_DELIM = "="
)

type MetaData struct {
	Title  string
	Author string
	Others map[string]string
}

type BlogData struct {
	Meta       MetaData
	Origin     string
	Follows    map[string]string
	Posts      []BlogItem
	Pagination Pagination
}

type Pagination struct {
	Prev string
	Next string
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
	temp := strings.Split(line, CONFIG_DELIM)
	if len(temp) >= 2 {
		return strings.TrimSpace(temp[0]), strings.TrimSpace(strings.Join(temp[1:], "=")), true
	}

	return "", "", false
}

func popLines(slice []string, index int) []string {
	return slice[index:]
}

func parsePagination(item string) (bool, string, bool) {
	if strings.HasPrefix(item, PAGINATION_NEXT) {
		return true, strings.TrimSpace(strings.TrimPrefix(item, PAGINATION_NEXT)), true
	} else if strings.HasPrefix(item, PAGINATION_PREV) {

		return false, strings.TrimSpace(strings.TrimPrefix(item, PAGINATION_PREV)), true
	}
	return false, "", false
}

func Parse(src string) {
	data := strings.Split(src, "\n")
	blg := BlogData{
		Meta:       MetaData{Others: make(map[string]string)},
		Origin:     "",
		Follows:    make(map[string]string),
		Posts:      make([]BlogItem, 0),
		Pagination: Pagination{},
	}
	lastindex := 0
	//params parsing
	for _, item := range data {
		if strings.HasPrefix(item, COMMENT) {
			continue
		}
		if strings.HasPrefix(item, FOLLOW_BLOCK) {
			break
		}

		//fmt.Printf("%d |%s\n", index, item)
		if key, value, ok := readOption(item); ok {

			switch key {
			case "title":
				blg.Meta.Title = value
			case "author":
				blg.Meta.Author = value
			case "origin":
				blg.Origin = value
			default:
				blg.Meta.Others[key] = value
			}
			lastindex += 1
		}

	}
	data = popLines(data, lastindex)
	isInsideFollowList := false
	for index, item := range data {
		if strings.HasPrefix(item, FOLLOW_BLOCK) {
			isInsideFollowList = !isInsideFollowList
			if !isInsideFollowList {
				lastindex = index + 1
				break
			}
		}

		if isInsideFollowList {
			if k, v, ok := readOption(item); ok {
				if strings.HasPrefix(k, FOLLOW) {
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
	startedPostBlock := false
	hasNext, HasPrev := false, false
	for _, item := range data {
		item = strings.TrimSpace(item)

		if (strings.HasPrefix(item, PAGINATION_NEXT) || strings.HasPrefix(item, PAGINATION_PREV)) && !startedPostBlock {
			if isNext, url, ok := parsePagination(item); ok {
				if isNext && !hasNext {
					blg.Pagination.Next = url
					hasNext = true
				} else if !HasPrev {
					blg.Pagination.Prev = url
					HasPrev = true
				}
			}
			continue
		}

		if strings.HasPrefix(item, DATE_START) { //Read Dates
			if strings.HasSuffix(item, DATA_END) {
				startedPostBlock = true
				dt := strings.TrimPrefix(item, DATE_START)
				dt = strings.TrimSuffix(dt, DATA_END)
				tempItem.Date = dt
				continue
			}
		}

		if !strings.HasPrefix(item, POST_END) { // This Means post has not yet ended

			if strings.HasPrefix(item, POST_META) && !blogMetaLock { //Check for post metas
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

		} else if strings.HasPrefix(item, POST_END) {
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
