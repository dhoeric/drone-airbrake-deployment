package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

// build number set at compile-time
var build = "0"

// Version set at compile-time
var Version string

func main() {
	if Version == "" {
		Version = fmt.Sprintf("0.0.1+%s", build)
	}

	app := cli.NewApp()
	app.Name = "Airbrake deploy notifier"
	app.Usage = "Notify Airbrake on new deployment"
	app.Copyright = "Copyright (c) 2018 Eric Ho"
	app.Authors = []cli.Author{
		{
			Name:  "Eric Ho",
			Email: "dho.eric@gmail.com",
		},
	}
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "airbrake-project-id",
			Usage:  "Airbrake project ID",
			EnvVar: "AIRBRAKE_PROJECT_ID,PLUGIN_AIRBRAKE_PROJECT_ID",
		},
		cli.StringFlag{
			Name:   "airbrake-project-key",
			Usage:  "Airbrake project key",
			EnvVar: "AIRBRAKE_PROJECT_KEY,PLUGIN_AIRBRAKE_PROJECT_KEY",
		},
		cli.StringFlag{
			Name:   "airbrake-environment",
			Usage:  "Deployed environment in Drone",
			EnvVar: "PLUGIN_AIRBRAKE_ENVIRONMENT",
		},
		cli.StringFlag{
			Name:   "build-author",
			Usage:  "Drone build author",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "build-commit",
			Usage:  "Drone build commit SHA",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "repo-link",
			Usage:  "Link of repository",
			EnvVar: "DRONE_REPO_LINK",
		},
	}

	// Override a template
	cli.AppHelpTemplate = `
NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}
REPOSITORY:
    Github: https://github.com/dhoeric/drone-airbrake-deployment
`

	if err := app.Run(os.Args); err != nil {
		fmt.Println("drone-airbrake-deployment Error: ", err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		Config: Config{
			ProjectID:   c.String("airbrake-project-id"),
			ProjectKey:  c.String("airbrake-project-key"),
			Environment: c.String("airbrake-environment"),
			BuildAuthor: c.String("build-author"),
			BuildCommit: c.String("build-commit"),
			RepoLink:    c.String("repo-link"),
		},
	}

	return plugin.Exec()
}
