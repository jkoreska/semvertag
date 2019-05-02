package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func main() {

	action := ""
	which := ""

	if 3 == len(os.Args) {
		action = os.Args[1]
		which = os.Args[2]
	}

	switch action {
	case "set":
	case "get":

		repo, err := git.PlainOpen(".git")
		if nil != err {
			softPanic(err.Error())
		}

		version := getLatest(repo)

		switch which {
		case "next":
			version = version.Next()
		case "nextminor":
			version = version.NextMinor()
		case "nextmajor":
			version = version.NextMajor()
		}

		switch action {
		case "get":
			fmt.Print(version)
		case "set":
			fmt.Print(version)
			setVersion(repo, version)
		}

	default:
		softPanic(fmt.Sprintf("Usage: %s [get|set] [latest|next|nextminor|nextmajor]", os.Args[0]))
	}
}

func setVersion(repo *git.Repository, version Version) {

	ref, err := repo.Head()
	if nil != err {
		softPanic(err.Error())
	}

	_, err = repo.CreateTag(version.String(), ref.Hash(), nil)
	if nil != err {
		softPanic(err.Error())
	}
}

func getLatest(repo *git.Repository) Version {

	latest := Version{}

	for _, v := range getVersions(repo) {
		if v.Major > latest.Major {
			latest = v
		} else if v.Major == latest.Major {
			if v.Minor > latest.Minor {
				latest = v
			} else if v.Minor == latest.Minor {
				if v.Point > latest.Point {
					latest = v
				}
			}
		}
	}

	return latest
}

func getVersions(repo *git.Repository) []Version {

	tagrefs, err := repo.Tags()
	if nil != err {
		softPanic(err.Error())
	}

	versions := make([]Version, 0)
	versionRegex, _ := regexp.Compile("^\\d+\\.\\d+\\.\\d+$")

	tagrefs.ForEach(func(ref *plumbing.Reference) error {

		tag := ref.Name().Short()

		if versionRegex.MatchString(tag) {

			versionParts := strings.Split(tag, ".")

			versions = append(versions, Version{
				Major: atoi(versionParts[0]),
				Minor: atoi(versionParts[1]),
				Point: atoi(versionParts[2]),
			})
		}

		return nil
	})

	return versions
}

func atoi(s string) uint64 {
	i, _ := strconv.ParseInt(s, 10, 0)
	return uint64(i)
}

func softPanic(s string) {
	fmt.Println(s)
	os.Exit(1)
}
