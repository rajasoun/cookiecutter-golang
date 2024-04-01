package testutils

import (
	"os"
	"strings"
)

// IsRunningAsGoTest checks whether the current process is running as a Go test.
func IsRunningAsGoTest() bool {
	testFlags := []string{"-test", "-test.run", "-test.v", "-test.timeout", "-test.bench"}
	for _, arg := range os.Args {
		for _, flag := range testFlags {
			if strings.HasPrefix(arg, flag) {
				return true
			}
		}
	}
	return false
}
