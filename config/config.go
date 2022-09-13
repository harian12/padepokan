package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"padepokan/command"
	"padepokan/logger"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		logger.ErrorFunc(err.Error(), err)
		log.Fatal("Error Load ENV File")
	}

	ConnectionDB()

	arg := flag.Arg(0)
	if arg != "" {
		command.InitCommands(ConnectionDB())
	}
}
