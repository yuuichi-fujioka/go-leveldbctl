package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/yuuichi-fujioka/go-leveldbctl/pkg/leveldbctl"
)

func kvfmt(ishex bool, kvarg string) ([]byte, string) {
	if !ishex {
		return []byte(kvarg), "%s"
	}
	kv, err := hex.DecodeString(kvarg)
	if err != nil {
		log.Fatal(err)
	}
	return kv, "%x"
}

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
		cli.BoolFlag{
			Name:  "hexkey, xk",
			Usage: "get / put hexadecimal keys",
		},
		cli.BoolFlag{
			Name:  "hexvalue, xv",
			Usage: "get / put hexadecimal values",
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
				key, kfmt := kvfmt(c.GlobalBool("xk"), c.Args()[0])
				value, vfmt := kvfmt(c.GlobalBool("xv"), c.Args()[1])
				err := leveldbctl.Put(c.GlobalString("dbdir"), key, value)
				if err != nil {
					return err
				}
				fmtstr := fmt.Sprintf("put %s: %s into %s.\n", kfmt, vfmt, "%s")
				fmt.Printf(fmtstr, key, value, c.GlobalString("dbdir"))
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
				key, _ := kvfmt(c.GlobalBool("xk"), c.Args()[0])
				value, ok, err := leveldbctl.Get(c.GlobalString("dbdir"), key)
				if err != nil {
					return err
				}
				if !ok {
					return cli.NewExitError(fmt.Sprintf("%v is not found.\n", key), 1)
				}

				fmt.Println(value)
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
				key, kfmt := kvfmt(c.GlobalBool("xk"), c.Args()[0])
				err := leveldbctl.Delete(c.GlobalString("dbdir"), key)
				if err != nil {
					return err
				}
				fmtstr := fmt.Sprintf("%s is deleted\n", kfmt)
				fmt.Printf(fmtstr, key)
				return nil
			},
		},
		{
			Name:      "search",
			Aliases:   []string{"s"},
			Usage:     "Search key prefix from a LevelDB",
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
				key, _ := kvfmt(c.GlobalBool("xk"), c.Args()[0])
				value, ok, err := leveldbctl.Search(c.GlobalString("dbdir"), key)
				if err != nil {
					return err
				}
				if !ok {
					return cli.NewExitError(fmt.Sprintf("%v is not found.\n", key), 1)
				}

				fmt.Println(value)
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
