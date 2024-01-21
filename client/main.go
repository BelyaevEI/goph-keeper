//go:build linux || windows || darwin

package main

import (
	"github.com/BelyaevEI/GophKeeper/client/cmd"
)

var (
	Version   string
	BuildTime string
)

func main() {
	cmd.Version = Version
	cmd.BuildTime = BuildTime
	cmd.Execute()
}
