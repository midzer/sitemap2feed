package main

import (
    "os"

    "github.com/midzer/go-cli-template/app"
)
func main() {
    os.Exit(app.CLI(os.Args[1:]))
}
