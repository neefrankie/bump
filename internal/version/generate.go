package main

import (
	"flag"
	"fmt"
	"github.com/neefrankie/bump/cmd"
	"github.com/neefrankie/bump/pkg/semver"
	"io/ioutil"
	"log"
)

var (
	major bool
	minor bool
	patch bool
)

var tmpl = `package %s

var Version = "%s"
`

func main() {
	flag.BoolVar(&major, "major", false, "Increase major version")
	flag.BoolVar(&minor, "minor", false, "Increase minor version")
	flag.BoolVar(&patch, "patch", false, "Increase patch version")

	flag.Parse()

	var sv semver.SemVer
	var err error
	switch {
	case major:
		sv, err = cmd.Incr(semver.VerPartMajor, "", true)
	case minor:
		sv, err = cmd.Incr(semver.VerPartMinor, "", true)
	case patch:
		sv, err = cmd.Incr(semver.VerPartPatch, "", true)
	}

	if err != nil {
		log.Fatal(err)
	}

	content := fmt.Sprintf(tmpl, "cmd", sv.String())

	log.Printf("Generated content %s", content)

	err = ioutil.WriteFile("cmd/version.go", []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("build/version.txt", []byte(sv.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
