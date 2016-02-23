package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func getRemote() string {
    var (
        out []byte
        err error
    )

    cmd := "git"
    args := []string {"remote", "get-url", "origin", "--push"}
    if out, err = exec.Command(cmd, args...).Output(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    return strings.TrimSpace(string(out))
}

func getBranch() string {
    var (
        out []byte
        err error
    )

    cmd := "git"
    args := []string {"rev-parse", "--abbrev-ref", "HEAD"}
    if out, err = exec.Command(cmd, args...).Output(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    return strings.TrimSpace(string(out))
}

// path: username/repo
func usernameAndRepo(path string) (username, repo string) {
    slices := strings.Split(path, "/")

    username = slices[0]

    repo = strings.TrimRight(slices[1], ".git")

    repo = strings.TrimSpace(repo)

    return
}

func parseTarget(target string, currentRepo *Repo) *Repo {
    var branch, username, repo string

    // target-user/target-repo@target-branch
    if strings.Contains(target, "@") && strings.Contains(target, "/") {
        slices := strings.Split(target, "@")

        path := slices[0]

        branch = strings.TrimSpace(slices[1])

        username, repo = usernameAndRepo(path)

    // target-user/target-repo
    } else if strings.Contains(target, "/") {
        branch = "master"

        username, repo = usernameAndRepo(target)

    // @target-branch
    } else if strings.Contains(target, "@") {
        username = currentRepo.username

        branch = strings.TrimLeft(target, "@")

        repo = currentRepo.name

    // target-user
    } else {
        username = strings.TrimSpace(target)

        branch = "master"

        repo = currentRepo.name
    }

    return &Repo {
        remote: "",

        username: username,

        name: repo,

        branch: branch,
    }
}


func parseGithubRemote(remote string) (string, string) {
    var path string

    if strings.Contains(remote, "git@github.com:") {
        path = strings.Replace(remote, "git@github.com:", "", 1)
    } else {
        path = strings.Replace(remote, "https://github.com/", "", 1)
    }

    username, repo := usernameAndRepo(path)

    return username, repo
}
