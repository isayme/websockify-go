package websockify

import "fmt"

// version info
var (
	Name        = "@isayme/websockify"
	Version     = "unknown"
	BuildTime   = "unknown"
	GitRevision = "unknown"
)

// PrintVersion print version
func PrintVersion() {
	fmt.Printf("name: %s\n", Name)
	fmt.Printf("version: %s\n", Version)
	fmt.Printf("buildTime: %s\n", BuildTime)
	fmt.Printf("gitRevision: %s\n", GitRevision)
}
