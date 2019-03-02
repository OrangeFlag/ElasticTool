package main

import (
	"context"
	"fmt"
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
				err = cli.ShowSubcommandHelp(c)
				return err
			},
			Subcommands: []cli.Command{
				{
					Name:  "indices",
					Usage: "show indices",
					Before:  func(c *cli.Context) error{
						c.App.Metadata["indeces"], err = client.CatIndices().Columns("health,status,index,docs.count").Do(context.Background())
						return err
					},
					Action: func(c *cli.Context) error {
						indeces := c.App.Metadata["indeces"].(elastic.CatIndicesResponse)

						fmt.Println("Health | Status | Index | DocsCount")
						for _, row := range indeces{
							fmt.Println(row.Health, row.Status, row.Index, row.DocsCount)
						}
						return nil
					},
					Subcommands: []cli.Command{
						{
							Name:  "count",
							Usage: "show count of indices",
							Action: func(c *cli.Context) error {
								indeces := c.App.Metadata["indeces"].(elastic.CatIndicesResponse)

								fmt.Println(len(indeces))
								return nil
							},
						},
					},
				},
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		client, err = elastic.NewClient(elastic.SetURL("http://" + host + ":" + port))
		return err

	}

	app.Action = func(c *cli.Context) error {
		if client != nil{
			fmt.Println("Server is available")
		}
		return nil
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
