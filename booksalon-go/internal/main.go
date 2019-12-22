package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

const version string = "v0.1.0"

var userSessions = make(map[string]string, 10)

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
	initDB()
	defer db.Close()

	r := gin.Default()
	initView(r)
	r.Run()
}
