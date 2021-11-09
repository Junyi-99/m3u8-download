package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "uri",
				Aliases:  []string{"u"},
				Usage:    "The uri of m3u8 file",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "video-name",
				Aliases:  []string{"n"},
				Usage:    "The output video name",
				Required: false,
				Value:    "output.mp4",
			},
			&cli.IntFlag{
				Name:     "thread",
				Aliases:  []string{"t"},
				Usage:    "concurrent thread",
				Required: false,
				Value:    8,
			},
			&cli.BoolFlag{
				Name:     "https",
				Aliases:  []string{"s"},
				Usage:    "Force HTTPS",
				Required: false,
				Value:    true,
			},
			&cli.StringFlag{
				Name:     "cookies",
				Usage:    "Custom Cookies",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "headers",
				Usage:    "Custom Headers",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "user-agent",
				Aliases:  []string{"ua"},
				Usage:    "Custom User-Agent",
				Required: false,
			},
		},
		Name:  "m3u8-download",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
