package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/bored-engineer/sitemap"
	JSON "github.com/mmcdole/gofeed/json"
)

func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	
	return 0
}

type appEnv struct {
}

func (app *appEnv) fromArgs(args []string) error {
	return nil
}

func (app *appEnv) run() error {

	urls, err := sitemap.Fetch(context.TODO(), "https://sitemaps.org/sitemap.xml")
	if err != nil {
		panic(err)
	}
	
	var items []*JSON.Item
	for _, url := range urls {
		log.Println(url.LastModification, url.Location)
		item := &JSON.Item {
			ID: url.Location,
			URL: url.Location,
			DatePublished: url.LastModification.String(),
		}
		items = append(items, item)
	}

	data := &JSON.Feed {
		Version:    "https://jsonfeed.org/version/1.1",
		Title:      "Sitemaps",
		HomePageURL: "https://sitemaps.org/",
		FeedURL:    "https://sitemaps.org/feed.json",
		Items: items,
	}

	result, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("feed.json")
	if err != nil{
		panic(err)
	}
	defer f.Close()

	f.Write(result)

	return nil
}
