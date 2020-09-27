package cmd

import (
	"fmt"
	"github.com/neefrankie/bump/pkg/semver"
	"os/exec"
	"strings"
)

func GetLatestTag() (string, error) {
	b, err := exec.Command("git", "tag", "-l", "--sort=v:refname").Output()

	if err != nil {
		return "", err
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
		return semver.SemVer{}, err
	}

	if vStr == "" {
		return semver.SemVer{}, nil
	}

	return semver.Parse(vStr)
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
	current, err := LatestVersion()
	if err != nil {
		fmt.Printf("Error finding latest version tag: %v\n", err)
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
	} else {
		err := AddTag(current, m)
		if err != nil {
			fmt.Printf("Error adding tag: %v\n", err)
			return semver.SemVer{}, err
		} else {
			fmt.Printf("Version upgraded %s -> %s\n", current, newVer)
		}
	}

	return newVer, nil
}
