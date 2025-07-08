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
				Action: func(ctx context.Context, cmd *cli.Command) error {
					scaffolder.Scaffold([]string{".gitignore", ".pre-commit-config.yaml", "Taskfile.yaml", "main_test.go"}, "")

					return nil
				},
			},
			{
				Name:  "moon",
				Usage: "Set up moonrepo",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					scaffolder.Scaffold([]string{".moon/toolchain.yml", ".moon/workspace.yml", ".moon/tasks.yml"}, "")

					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
