package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/yuuichi-fujioka/go-leveldbctl/pkg/leveldbctl"
)

func main() {
	app := cli.NewApp()
	app.Name = "leveldbctl"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "dbdir, d",
			Value:  "./",
			Usage:  "LevelDB Directory",
			EnvVar: "LEVELDB_DIR",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialize a LevelDB",
			Action: func(c *cli.Context) error {
				err := leveldbctl.Init(c.GlobalString("dbdir"))
				if err != nil {
					return err
				}
				fmt.Printf("%s is initialized as LevelDB\n", c.GlobalString("dbdir"))
				return nil
			},
		},
		{
			Name:    "walk",
			Aliases: []string{"w"},
			Usage:   "Walk in a LevelDB",
			Action: func(c *cli.Context) error {
				err := leveldbctl.Walk(c.GlobalString("dbdir"), func(k, v string) {
					fmt.Printf("%s: %s\n", k, v)
				})
				return err
			},
		},
		{
			Name:    "keys",
			Aliases: []string{"k"},
			Usage:   "Search all keys in a LevelDB",
			Action: func(c *cli.Context) error {
				err := leveldbctl.Walk(c.GlobalString("dbdir"), func(k, _ string) {
					fmt.Printf("%s\n", k)
				})
				return err
			},
		},
		{
			Name:      "put",
			Aliases:   []string{"p"},
			Usage:     "Put a value into a LevelDB",
			ArgsUsage: "key value",
			Action: func(c *cli.Context) error {
				if c.NArg() != 2 {
					if c.NArg() < 2 {
						fmt.Println("[ERROR] key and value are required.")
					}
					if c.NArg() > 2 {
						fmt.Println("[ERROR] Many arguments are passed.")
					}
					return cli.ShowSubcommandHelp(c)
				}
				key := c.Args()[0]
				value := c.Args()[1]
				err := leveldbctl.Put(c.GlobalString("dbdir"), key, value)
				if err != nil {
					return err
				}
				fmt.Printf("put %s: %s into %s.\n", key, value, c.GlobalString("dbdir"))
				return nil
			},
		},
		{
			Name:      "get",
			Aliases:   []string{"g"},
			Usage:     "Gut a value from a LevelDB",
			ArgsUsage: "key",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					if c.NArg() < 1 {
						fmt.Println("[ERROR] key is required.")
					}
					if c.NArg() > 1 {
						fmt.Println("[ERROR] Many arguments are passed.")
					}
					return cli.ShowSubcommandHelp(c)
				}
				key := c.Args()[0]
				value, ok, err := leveldbctl.Get(c.GlobalString("dbdir"), key)
				if err != nil {
					return err
				}
				if ok {
					fmt.Println(value)
				} else {
					fmt.Printf("%s is not found.\n", key)
				}

				return nil
			},
		},
		{
			Name:      "delete",
			Aliases:   []string{"d"},
			Usage:     "Delete a value from a LevelDB",
			ArgsUsage: "key",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					if c.NArg() < 1 {
						fmt.Println("[ERROR] key is required.")
					}
					if c.NArg() > 1 {
						fmt.Println("[ERROR] Many arguments are passed.")
					}
					return cli.ShowSubcommandHelp(c)
				}
				key := c.Args()[0]
				err := leveldbctl.Delete(c.GlobalString("dbdir"), key)
				if err != nil {
					return err
				}
				fmt.Printf("%s is deleted\n", key)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
