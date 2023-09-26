package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bored-engineer/sitemap"
	"github.com/gorilla/feeds"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func writeFile(fname string, c func() (string, error)) error {
	f, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf("can't create %s: %w", fname, err)
	}
	defer f.Close()

	s, err := c()
	if err != nil {
		return fmt.Errorf("can't convert feed: %w", err)
	}

	_, err = f.WriteString(s)
	if err != nil {
		return fmt.Errorf("can't write to %s: %w", fname, err)
	}

	return nil
}

func run() error {
	urls, err := sitemap.Fetch(context.Background(), "https://www.asheeshkg.com/post-sitemap.xml")
	if err != nil {
		return fmt.Errorf("can't fetch sitemap: %w", err)
	}

	feed := &feeds.Feed{
		Title: "Sitemaps",
		Link:  &feeds.Link{Href: "https://www.asheeshkg.com/"},
	}

	for _, url := range urls {
		log.Println(url.LastModification, url.Location)
		updated, err := time.Parse("2006-01-02", url.LastModification.String())
		if err != nil {
			return fmt.Errorf("can't parse last modification date: %w", err)
		}

		item := &feeds.Item{
			Id:      url.Location,
			Link:    &feeds.Link{Href: url.Location},
			Updated: updated,
		}
		feed.Add(item)
	}

	// write atom.xml
	if err := writeFile("atom.xml", feed.ToAtom); err != nil {
		return err
	}

	// write rss.xml
	if err := writeFile("rss.xml", feed.ToRss); err != nil {
		return err
	}

	// write feed.json
	if err := writeFile("feed.json", feed.ToJSON); err != nil {
		return err
	}

	return nil
}
