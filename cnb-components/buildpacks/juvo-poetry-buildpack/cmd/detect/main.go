package main

import (
	juvo "github.com/juris-futura/juvo-poetry-buildpack/juvo"
	"github.com/paketo-buildpacks/packit"
)

func main() {
	packit.Detect(juvo.Detect(juvo.PhysicalFs{}))
}
