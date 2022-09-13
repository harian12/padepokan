package command

import (
	"flag"
	"github.com/urfave/cli"
	"gorm.io/gorm"
	"log"
	"os"
	"padepokan/command/cmd_variabel"
	"padepokan/command/src_command"
	"padepokan/database/migrations"
	"padepokan/database/seeders"
)

func InitCommands(db *gorm.DB) {
	arg := flag.Arg(1)

	app_name_cmd := os.Getenv("APP_NAME_CMD")

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				migrations.DBMigrate(db)
				return nil
			},
		},
		{
			Name: "db:seeder",
			Action: func(c *cli.Context) error {
				seeders.DBSeed(db)
				return nil
			},
		},
		{
			Name: "make:src",
			Action: func(c *cli.Context) error {
				src_command.MakeSrcRepository(arg)
				src_command.MakeSrcService(arg, app_name_cmd)
				src_command.MakeSrcController(arg, app_name_cmd)
				cmd_variabel.AddSrcVariabel(arg)
				return nil
			},
		},
		{
			Name: "make:c",
			Action: func(c *cli.Context) error {
				src_command.MakeSrcController(arg, app_name_cmd)
				cmd_variabel.AddControllerVariabel(arg)
				return nil
			},
		},
	}
	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
