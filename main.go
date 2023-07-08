package main

import (
    "os"

    "github.com/midzer/sitemap2feed/app"
)
func main() {
    os.Exit(app.CLI(os.Args[1:]))
}
