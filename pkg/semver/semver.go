package semver

import (
	"errors"
	"strconv"
	"strings"
)

// SemVer represents a parsed value of semantic version string.
type SemVer struct {
	Major int
	Minor int
	Patch int
}

// Parse parses a string into to SemVer
func Parse(v string) (SemVer, error) {
	if len(v) == 0 {

		return SemVer{}, errors.New("empty version string")
	}

	var parts = make([]int, 0)

	v = strings.TrimPrefix(v, "v")

	for _, v := range strings.Split(v, ".") {
		n, err := strconv.Atoi(v)
		if err != nil {
			return SemVer{}, err
		}

		parts = append(parts, n)
	}

	gap := len(parts) - 3
	if gap > 0 {
		for i := 0; i < gap; i++ {
			parts = append(parts, 0)
		}
	}

	return SemVer{
		Major: parts[0],
		Minor: parts[1],
		Patch: parts[2],
	}, nil
}

func (s SemVer) String() string {
	var buf strings.Builder
	buf.WriteString("v")
	buf.WriteString(strconv.Itoa(s.Major))
	buf.WriteString(".")
	buf.WriteString(strconv.Itoa(s.Minor))
	buf.WriteString(".")
	buf.WriteString(strconv.Itoa(s.Patch))

	return buf.String()
}

// Compare compares two semantic versions.
// Returns negative number if s should come before (is smaller) other;
// 0 if the two are equal;
// Positive number if s should come after (is larger than) other.
func (s SemVer) Compare(other SemVer) int {
	diff := s.Major - other.Major
	if diff != 0 {
		return diff
	}

	diff = s.Minor - other.Minor
	if diff != 0 {
		return diff
	}

	return s.Patch - other.Patch
}

// Equal tests if two semantic version string are the same.
func (s SemVer) Equal(other SemVer) bool {
	return s.Compare(other) == 0
}

// LargeThan tests whether semantic version a is larger than b
func (s SemVer) LargeThan(other SemVer) bool {
	return s.Compare(other) > 0
}

// SmallThan tests whether semantic version a is smaller than b.
func (s SemVer) SmallThan(other SemVer) bool {
	return s.Compare(other) < 0
}

func (s SemVer) IncrMajor() SemVer {
	s.Major++
	s.Minor = 0
	s.Patch = 0

	return s
}

func (s SemVer) IncrMinor() SemVer {
	s.Minor++
	s.Patch = 0

	return s
}

func (s SemVer) IncrPatch() SemVer {
	s.Patch++

	return s
}
