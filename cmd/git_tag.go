package cmd

import (
	"errors"
	"fmt"
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
		return "", errors.New("no version tag set for current repo")
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

func AddTag(v SemVer, m string) (string, error) {
	args := make([]string, 0)

	if m != "" {
		m = fmt.Sprintf(m, v)
		args = []string{"tag", "-a", v.String(), "-m", m}
	} else {
		args = []string{"tag", v.String()}
	}

	b, err := exec.Command("git", args...).Output()
	if err != nil {
		return "", err
	}

	return string(b), nil
}
