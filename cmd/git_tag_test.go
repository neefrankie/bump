package cmd

import (
	"github.com/neefrankie/bump/pkg/semver"
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

func TestIncr(t *testing.T) {
	type args struct {
		p      semver.VerPart
		m      string
		dryRun bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Dry run",
			args: args{
				p:      semver.VerPartMajor,
				m:      "",
				dryRun: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Incr(tt.args.p, tt.args.m, tt.args.dryRun)
			if (err != nil) != tt.wantErr {
				t.Errorf("Incr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("%s", got)
		})
	}
}

func TestIsClean(t *testing.T) {
	ok, err := IsClean()
	if err != nil {
		t.Error(err)
	}

	t.Logf("%t", ok)
}
