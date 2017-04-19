// Package EDDNClient provides an interface to EDDN.  Currently it only
// provides subscriber support, but other features will be added in the future.
package EDDNClient

// EddnAddress is a simple constant for the ZeroMQ relay used by EDDN.
const EddnAddress = "tcp://eddn-relay.elite-markets.net:9500"

// version contains the current version in the form major, minor, and revision.
var version = [...]int{0, 0, 1}

// Version is a simple function that returns the major, minor, and revision
// versions.
func Version() (major int, minor int, revision int) {
	return version[0], version[1], version[2]
}
