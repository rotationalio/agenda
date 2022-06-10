package main

import (
	"os"

	"github.com/rotationalio/agenda/pkg"
	"github.com/rotationalio/agenda/pkg/server"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:    "agenda",
		Version: pkg.Version(),
		Commands: []*cli.Command{
			{
				Name:   "serve",
				Usage:  "serve the agenda server on port 55108",
				Action: serve,
				Flags:  []cli.Flag{},
			},
		},
	}

	app.Run(os.Args)
}

func serve(c *cli.Context) error {
	srv, err := server.New()
	if err != nil {
		return cli.Exit(err, 1)
	}

	if err = srv.Serve(":55108"); err != nil {
		return cli.Exit(err, 1)
	}
	return nil
}
