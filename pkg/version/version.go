package version

import (
	"encoding/json"
	"fmt"
	"runtime"
)

var (
	version   string
	gitBranch string
	gitTag    string
	gitCommit string
	buildDate string
)

type Info struct {
	Version   string `json:"version"`
	GitBranch string `json:"gitBranch"`
	GitTag    string `json:"gitTag"`
	GitCommit string `json:"gitCommit"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
	Compiler  string `json:"compiler"`
	Platform  string `json:"platform"`
}

func (v Info) String() string {
	j, _ := json.Marshal(&v)
	return string(j)
}

func GetVersion() Info {
	return Info{
		Version:   version,
		GitBranch: gitBranch,
		GitTag:    gitTag,
		GitCommit: gitCommit,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
