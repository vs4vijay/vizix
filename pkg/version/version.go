package version

import "fmt"

type Version struct {
	Version   string `json:"version"`
	GitCommit string `json:"git-commit"`
	BuildTime string `json:"build-time"`
}

func (v *Version) String() string {
	return v.Version
}

func (v *Version) FormatVersion(short bool) string {
	if short {
		return fmt.Sprintf("%s+g%s", v.Version, v.GitCommit[:7])
	}
	return fmt.Sprintf("%#v", v)
}

var (
	SemVer    = "canary"
	GitCommit = ""
	BuildTime = ""
)

func getVersion() string {
	if BuildTime == "" {
		return SemVer
	}
	return SemVer + "+" + BuildTime
}

func New() *Version {
	return &Version{
		Version:   getVersion(),
		GitCommit: GitCommit,
	}
}
