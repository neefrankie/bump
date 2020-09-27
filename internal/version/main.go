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
	major   bool
	minor   bool
	patch   bool
	message string
)

var tmpl = `// Code generated .* DO NOT EDIT.

package %s

var Version = "%s"
`

func main() {
	flag.BoolVar(&major, "major", false, "Increase major version")
	flag.BoolVar(&minor, "minor", false, "Increase minor version")
	flag.BoolVar(&patch, "patch", false, "Increase patch version")
	flag.StringVar(&message, "message", "", "Message for annotated tag")

	flag.Parse()

	var part semver.VerPart
	switch {
	case major:
		part = semver.VerPartMajor
	case minor:
		part = semver.VerPartMinor
	case patch:
		part = semver.VerPartPatch
	default:
		log.Fatal("one the flags -major, -minor or -patch is required")
	}

	sv, err := cmd.Incr(part, "", true)

	if err != nil {
		log.Fatal(err)
	}

	content := fmt.Sprintf(tmpl, "cmd", sv.String())

	log.Printf("Generated content %s", content)

	err = ioutil.WriteFile("cmd/version.go", []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}

	ok, err := cmd.IsClean()
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		err := cmd.Commit("cmd/version.go", fmt.Sprintf("Bump version %s", sv.String()))
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = cmd.Incr(part, message, false)
	if err != nil {
		log.Fatal()
	}
}
