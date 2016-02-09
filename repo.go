package main

import (
    "strings"
)

type Repo struct {
    remote string

    username string

    name string

    branch string
}

func NewRepo(remote string, branch string) *Repo {

    username, repo := parseGithubRemote(remote)

    return &Repo {
        remote: remote,

        username: username,

        name: repo,

        branch: branch,
    }
}

func (r Repo) IsGithub() bool {
    return strings.Contains(r.remote, "github.com")
}
