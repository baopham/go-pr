package main

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/toqueteos/webbrowser"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-pr"
	app.Usage = "Create a Pull Request for the current branch"
	app.UsageText = `
    go-pr @target-branch # default target-user to your Github user
    go-pr target-user/target-repo@target-branch
    go-pr target-user/target-repo # default target-branch to master
    go-pr target-user # default target-branch to master and target-repo to the name of your Github repo
    `
	app.Version = "0.0.2"
	app.Author = "Bao Pham"
	app.Email = "gbaopham@gmail.com"

	app.Before = func(c *cli.Context) error {
		if len(c.Args()) < 1 {
			return errors.New("Target is required")
		}

		return nil
	}

	app.Action = func(c *cli.Context) {
		remote := getRemote()

		branch := getBranch()

		currentRepo := NewRepo(remote, branch)

		target := parseTarget(c.Args().First(), currentRepo)

		if !currentRepo.IsGithub() {
			println("Only support Github for now")
			os.Exit(1)
		}

		webbrowser.Open(fmt.Sprintf("https://github.com/%s/%s/compare/%s...%s:%s",
			target.username, target.name, target.branch,
			currentRepo.username, currentRepo.branch))
	}

	app.Run(os.Args)
}
