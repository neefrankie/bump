package cmd

import (
	"errors"
	"fmt"
	"github.com/neefrankie/bump/pkg/semver"
	"log"
	"os/exec"
	"strings"
)

func GetLatestTag() (string, error) {
	b, err := exec.Command("git", "tag", "-l", "--sort=v:refname").Output()

	if err != nil {
		return "", fmt.Errorf("latest tag: " + err.Error())
	}
	tagStr := strings.TrimSpace(string(b))
	if tagStr == "" {
		return "", nil
	}

	tags := strings.Split(tagStr, "\n")

	var latest string
	for i := len(tags) - 1; i >= 0; i-- {
		latest = strings.TrimSpace(tags[i])

		if latest != "" {
			break
		}
	}

	return latest, nil
}

func LatestVersion() (semver.SemVer, error) {
	vStr, err := GetLatestTag()
	if err != nil {
		return semver.SemVer{}, fmt.Errorf("latest version: " + err.Error())
	}

	if vStr == "" {
		return semver.SemVer{}, nil
	}

	return semver.Parse(vStr)
}

func IsClean() (bool, error) {
	b, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return false, fmt.Errorf("is clean: %s", err)
	}

	log.Printf("%s\n", b)

	return string(b) == "", nil
}

func Commit(file string, m string) error {
	err := exec.Command("git", "add", file).Run()
	if err != nil {
		return err
	}

	return exec.Command("git", "commit", "-m", m).Run()
}

func AddTag(v semver.SemVer, m string) error {
	args := make([]string, 0)

	if m != "" {
		m = fmt.Sprintf(m, v)
		args = []string{"tag", "-a", v.String(), "-m", m}
	} else {
		args = []string{"tag", v.String()}
	}

	return exec.Command("git", args...).Run()
}

func Incr(p semver.VerPart, m string, dryRun bool) (semver.SemVer, error) {
	ok, err := IsClean()
	if err != nil {
		return semver.SemVer{}, fmt.Errorf("incr: %s", err)
	}

	if !ok {
		return semver.SemVer{}, errors.New("incr: repository not clean")
	}

	current, err := LatestVersion()
	if err != nil {
		return semver.SemVer{}, err
	}

	var newVer semver.SemVer
	switch p {
	case semver.VerPartMajor:
		newVer = current.IncrMajor()
	case semver.VerPartMinor:
		newVer = current.IncrMinor()
	case semver.VerPartPatch:
		newVer = current.IncrPatch()
	default:
		newVer = current
	}

	if dryRun {
		fmt.Printf("Version will be upgraded %s -> %s\n", current, newVer)
		return newVer, nil
	}

	err = AddTag(newVer, m)
	if err != nil {
		return semver.SemVer{}, fmt.Errorf("incr: error tagging: %s", err)
	}

	fmt.Printf("Version upgraded %s -> %s\n", current, newVer)

	return newVer, nil
}
