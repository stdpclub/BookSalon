package main

import (
	"BookSalon/booksalon-go/dbconn"
	"BookSalon/booksalon-go/router"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const version string = "v0.1.0"

func main() {
	app := &cli.App{
		Name:  "booksalon",
		Usage: "booksalon runserver",
		Action: func(c *cli.Context) error {
			initRun()
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func initRun() {
	db := dbconn.NewDBConn()
	defer db.Close()
	router.InitView().Run()
}
