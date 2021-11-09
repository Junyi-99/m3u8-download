package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
	"log"
	"m3u8-download/pool"
	"math/rand"
	"os"
	"time"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "uri",
				Aliases:  []string{"u"},
				Usage:    "The URL to the playlist.",
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
			PoolTest()
			fmt.Println("!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
func PoolTest() {
	b := mpb.New()
	p := pool.New(4, 100)
	var bars []*mpb.Bar

	for i := 0; i < 100; i++ {
		task := fmt.Sprintf("Thread #%02d:", i)
		p.Submit(func() {

			bar := b.AddBar(123*1024*1024,
				mpb.PrependDecorators(
					decor.Name(task, decor.WC{W: 25, C: decor.DidentRight}),
					decor.Name("downloading", decor.WCSyncSpaceR),
					decor.CountersKiloByte("%d / %d", decor.WCSyncWidth),
				),
				mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5})),
				mpb.BarRemoveOnComplete(),
			)
			bars = append(bars, bar)
			for !bar.Completed() {
				start := time.Now()
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				bar.IncrBy(9 * 1024 * 1024) // bytes

				bar.DecoratorEwmaUpdate(time.Since(start))
			}

		})
	}

	p.Wait()
}
