package main

import (
	"flag"
	"time"

	bond "github.com/kcmerrill/bond/james"
)

func main() {
	// setup our flags
	reload := flag.Duration("reload", 10*time.Second, "Reload configuration")
	config := flag.String("config", "bond.yml", "Action list filename")
	flag.Parse()

	// reload our config
	go bond.ReloadConfig(*config, *reload)

	// read in stdin
	bond.Read()
}
