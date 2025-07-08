package main

import (
	"context"
	"log"
	"os"

	"github.com/Rick-Phoenix/goutils/scaffolder"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := cli.Command{
		Name:  "scaffold",
		Usage: "Set up basic files for go/typescript projects",
		Commands: []*cli.Command{
			{
				Name:  "go",
				Usage: "Generate basic files for a go project.",
				Flags: []cli.Flag{
					dirFlag,
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					data := map[string]any{}

					return scaffolder.ScaffoldGo(c.String("dir"), data)
				},
			},
			{
				Name:  "pre-commit",
				Usage: "Generate a .pre-commit-config.yaml file",
				Flags: []cli.Flag{
					dirFlag,
					&cli.BoolFlag{
						Name:  "oxlint",
						Usage: "Add a hook for oxlint",
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					data := map[string]any{
						"Oxlint": c.Bool("oxlint"),
					}

					return scaffolder.ScaffoldPreCommit(c.String("dir"), data)
				},
			},
			{
				Name:  "moon",
				Usage: "Set up moonrepo",
				Flags: []cli.Flag{
					dirFlag,
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					data := map[string]any{}
					return scaffolder.ScaffoldMoonRepo(c.String("dir"), data)
				},
			},
			{
				Name:  "svelte",
				Usage: "Set up a svelte project",
				Flags: []cli.Flag{
					dirFlag,
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "The name of the project (to use in package.json)",
						Value:   "svelte-app",
					},
					&cli.BoolFlag{
						Name:    "wails",
						Aliases: []string{"w"},
						Usage:   "Whether the svelte app is part of a wails project",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					tmplData := map[string]any{
						"PackageName": cmd.String("name"),
						"IsWails":     cmd.Bool("wails"),
					}
					return scaffolder.ScaffoldSvelte(cmd.String("dir"), tmplData)
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

var dirFlag = &cli.StringFlag{
	Name:    "dir",
	Aliases: []string{"d"},
	Usage:   "The root directory for files generation.",
	Value:   ".",
}
