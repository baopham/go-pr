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

    username, repo := ParseGithubRemote(remote)

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

func ParseGithubRemote(remote string) (string, string) {
    var path string

    if strings.Contains(remote, "git@github.com:") {
        path = strings.Replace(remote, "git@github.com:", "", 1)
    } else {
        path = strings.Replace(remote, "https://github.com/", "", 1)
    }

    username, repo := usernameAndRepo(path)

    return username, repo
}
