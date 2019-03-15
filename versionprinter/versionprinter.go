package versionprinter

import "runtime"

// Version is an exported function
func Version() string {
	return runtime.Version()
}
