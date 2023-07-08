package main

import (
    "os"

    "github.com/midzer/sitemap2rss/app"
)
func main() {
    os.Exit(app.CLI(os.Args[1:]))
}
