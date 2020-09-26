package semver

type VerPart int

const (
	VerPartMajor VerPart = iota
	VerPartMinor
	VerPartPatch
)
