package main

import (
	"fmt"
	"os"

	"github.com/golangtokyo/gomopost"
	"github.com/urfave/cli"
)

// GlobalFlags mean flags specified by command line
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "a, address",
		Value: "http://localhost:8080",
		Usage: "specify server address (scheme and port required)",
	},
	cli.StringFlag{
		Name:  "n, name",
		Value: "anonymous",
		Usage: "specify name",
	},
	cli.StringFlag{
		Name:  "m, message",
		Value: "hello",
		Usage: "specify message",
	},
}

func defaultCommand(c *cli.Context) {
	client := gomopost.NewClient(c.GlobalString("address"))
	err := client.Post(c.GlobalString("name"), c.GlobalString("message"))
	if err != nil {
		fmt.Printf("failed to post: %s\n", err.Error())
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gomopost"
	app.Usage = "simple http client for gomobile instruction"
	app.Version = "0.1"
	app.Flags = GlobalFlags
	app.Action = defaultCommand

	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [options] [arguments...]

VERSION:
   {{.Version}}{{if or .Author .Email}}

AUTHOR:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
