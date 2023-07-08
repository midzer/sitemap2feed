package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bored-engineer/sitemap"
	"github.com/gorilla/feeds"
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
		log.Fatal(err)
	}

	feed := &feeds.Feed {
		Title:      "Sitemaps",
		Link: &feeds.Link { Href: "https://sitemaps.org/" },
	}

	for _, url := range urls {
		log.Println(url.LastModification, url.Location)
		updated, err := time.Parse("2006-01-02", url.LastModification.String())
		if err != nil {
			log.Fatal(err)
		}
		item := &feeds.Item {
			Id: url.Location,
			Link: &feeds.Link { Href: url.Location },
			Updated: updated,
		}
		feed.Add(item)
	}

	atom, err := feed.ToAtom()
    if err != nil {
        log.Fatal(err)
    }

	f, err := os.Create("atom.xml")
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString(atom)

    rss, err := feed.ToRss()
    if err != nil {
        log.Fatal(err)
    }

	f, err = os.Create("rss.xml")
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString(rss)

    json, err := feed.ToJSON()
    if err != nil {
        log.Fatal(err)
    }

	f, err = os.Create("feed.json")
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString(json)

	return nil
}
