package main

import (
	"learngo-pockets/logger/pocketlog"
	"os"
)

func main() {
    log := pocketlog.New(pocketlog.LevelDebug, os.Stderr)

    log.Debugf("Hello", "pish")
}