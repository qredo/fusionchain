package common

import (
	"fmt"
)

var (
	Version     = "v0.0.1" // Semantic version
	FullVersion = fmt.Sprintf("%s-%v", Version, CommitHash[0:8])
)
