package cmd

import (
	"testing"
)

func TestGetLatestTag(t *testing.T) {
	vStr, err := GetLatestTag()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf(vStr)
}

func TestAddTag(t *testing.T) {
	s, err := AddTag(NewPatch(), "Bump version %s")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(s)
}
