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

func Incr(p semver.VerPart, m string) {
	sv, err := LatestVersion()
	if err != nil {
		fmt.Printf("Error finding latest version tag: %v\n", err)
		return
	}

	fmt.Printf("Current version %s\n", sv)

	switch p {
	case semver.VerPartMajor:
		sv = sv.IncrMajor()
	case semver.VerPartMinor:
		sv = sv.IncrMinor()
	case semver.VerPartPatch:
		sv = sv.IncrPatch()
	}

	if err := AddTag(sv, m); err != nil {
		fmt.Printf("Error adding tag: %v\n", err)
		return
	}

	fmt.Printf("New version set to %s\n", sv)
	return
}
