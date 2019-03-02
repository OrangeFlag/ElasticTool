package main

import (
	"log"
	"os"

	"github.com/olivere/elastic"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ElasticTool"
	app.Usage = "a command line interface for managing a elasticsearch."

	var (
		host   string
		port   string
		client *elastic.Client
		err    error
	)

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "host, H",
			Value:       "127.0.0.1",
			Usage:       "Hostname or IP address.",
			Destination: &host,
		},
		cli.StringFlag{
			Name:        "port, p",
			Value:       "9200",
			Usage:       "Port number.",
			Destination: &port,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "cat",
			Usage: "cat API",
			Action: func(c *cli.Context) error {
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:  "indices",
					Usage: "cat indices",
					Action: func(c *cli.Context) error {
						return nil //TODO
					},
					Subcommands: []cli.Command{
						{
							Name:  "count",
							Usage: "count of indices",
							Action: func(c *cli.Context) error {
								return nil //TODO
							},
						},
					},
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "all, a",
							Usage: "include system indexes",
						},
					},
				},
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		//fmt.Println(host + ":" + port)

		client, err = elastic.NewClient(elastic.SetURL("http://" + host + ":" + port))
		if err != nil {
			log.Fatal(err)
		}

		//indeces, err := client.CatIndices().Do(context.Background())
		//fmt.Println(indeces)
		//
		//health, err := client.ClusterHealth().Do(context.Background())
		//
		//if err != nil || health == nil {
		//	return err
		//}
		//
		//fmt.Println(health)

		return nil

	}

	app.Action = func(c *cli.Context) error {
		return nil //TODO
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
