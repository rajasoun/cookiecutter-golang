package version

import (
	"fmt"
	"runtime"
)

// AppName is the name of the application, set at build time.
var AppName string

// Version is the main version number of the application.
var Version = "0.1.0"

// BuildDate is the date when the binary was built, set at build time.
var BuildDate string

// GitCommit is the git commit hash at the time of the build, set at build time.
var GitCommit string

// GoVersion is the version of the Go runtime used to compile the binary.
var GoVersion = runtime.Version()

// OsArch is the operating system and architecture used to build the binary.
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
