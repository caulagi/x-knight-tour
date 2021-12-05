package main

import (
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func userInput() (int, int, int) {
	var startX, startY, boardSize int

	app := &cli.App{
		Usage:       "a modified knight's tour",
		UsageText:   "x-knight-tour [global options] startX startY",
		Description: "uses warnsdorff algorithm to find a modified knight tour",
		HideHelp:    true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "help, h",
				Usage: "show help",
			},
			&cli.IntFlag{
				Name:        "boardSize",
				Value:       10,
				Usage:       "size of board",
				Destination: &boardSize,
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("help") {
				cli.ShowAppHelpAndExit(c, 0)
			}
			if c.NArg() > 0 {
				startX, _ = strconv.Atoi(c.Args().Get(0))
			}
			if c.NArg() > 1 {
				startY, _ = strconv.Atoi(c.Args().Get(1))
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	return startX, startY, boardSize
}
