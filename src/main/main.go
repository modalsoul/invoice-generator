package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	. "strings"
	"flag"
	"encoding/json"
	"log"
)

type Item struct {
	ID string
	JpName string
	EnName string
}

func (item Item) String() string {
	str := fmt.Sprintf("%s:%s:%s", item.ID, item.JpName, item.EnName)
	return str
}

func NewItem(id, jpName string) *Item {
	item := &Item{
		ID: id,
		JpName: jpName,
	}
	return item
}

func GetItems(url string)[]Item {
	var items []Item
	doc, _ := goquery.NewDocument(url)
	doc.Find(".a-fixed-left-grid   .a-spacing-large").Each(func(_ int, s *goquery.Selection) {
		rawId, _ := s.Attr("id")
		id := Trim(rawId, " ")
		if id != "" {
			anker := s.Find("a")
			href, _ := anker.Attr("href")
			rawTitle, _ := anker.Attr("title")
			title := Replace(rawTitle, "\u200b", "", -1)
			if href != "" && title != "" {
				splited := Split(href, "/")
				if len(splited) > 2 {
					id := splited[2]
					item := NewItem(id, title)
					items = append(items, *item)
				}
			}
		}
	})
	return items
}

func GetEnName(url string)string {
	var title string
	doc, _ := goquery.NewDocument(url)
	doc.Find("#productTitle").Each(func(_ int, s *goquery.Selection){
		title = s.Text()
	})
	return title
}

func TranslateItem(item Item)Item {
	const url = "http://www.amazon.com/gp/product/"
	enName := GetEnName(url + item.ID)
	item.EnName = enName
	return item
}

func main() {
	var wishlist = flag.String("wishlist", "", "target wishlist id")
	flag.Parse()
	url := "http://www.amazon.co.jp/gp/registry/wishlist/" + *wishlist + "/"
	for _, item := range GetItems(url) {
		json, err := json.Marshal(TranslateItem(item))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(json))
	}
}