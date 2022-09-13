package main

import (
	"flag"
	"padepokan/config"
	"padepokan/router"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		config.Run()
	} else {
		config.Run()
		router.AppRoute()
	}
}
