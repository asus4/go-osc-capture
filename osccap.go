package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	// Parse command args
	app.Name = "osccap"
	app.Usage = "-port 8000 -record osc.csv"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		// osccap record
		{
			Name:    "record",
			Aliases: []string{"rec", "r"},
			Usage:   "Record OSC to the file",
			Flags: []cli.Flag{
				cli.UintFlag{
					Name:  "port, p",
					Usage: "The listen port",
					Value: 8000,
				},
				cli.StringFlag{
					Name:  "multicast, m",
					Usage: "The subsclibe multicast address",
				},
			},
			Action: func(c *cli.Context) error {
				file := c.Args().First()
				if file == "" {
					return cli.NewExitError("Set the output file", 1)
				}
				err := recorder(c.Uint("port"), file)
				if err != nil {
					return err
				}
				return nil
			},
		},
		// osccap play
		{
			Name:    "play",
			Aliases: []string{"p"},
			Usage:   "Playback OSC in the file",
			Flags: []cli.Flag{
				cli.UintFlag{
					Name:  "port, p",
					Usage: "The listen port",
					Value: 8000,
				},
				cli.StringFlag{
					Name:  "addr, ip",
					Usage: "The target ip addresss",
					Value: "localhost",
				},
			},
			Action: func(c *cli.Context) error {
				file := c.Args().First()
				if file == "" {
					return cli.NewExitError("Set the playback file", 1)
				}
				err := player(c.String("addr"), c.Uint("port"), file)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}

	// run cli
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
