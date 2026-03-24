//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"

	// mage:import
	build "github.com/grafana/grafana-plugin-sdk-go/build"
)

// Includes linux/s390x and linux/ppc64le.
func BuildAllPlatforms() {
	b := build.Build{}
	mg.Deps(
		b.Linux,
		b.LinuxARM64,
		b.Windows,
		b.Darwin,
		b.DarwinARM64,
		buildLinuxS390X,
		buildLinuxPPC64LE,
	)
}

func buildLinuxS390X() error {
	return build.Build{}.Custom("linux", "s390x")
}

func buildLinuxPPC64LE() error {
	return build.Build{}.Custom("linux", "ppc64le")
}

var Default = BuildAllPlatforms
