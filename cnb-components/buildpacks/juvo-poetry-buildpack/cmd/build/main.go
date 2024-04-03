package main

import (
	"github.com/juris-futura/juvo-poetry-buildpack/poetry"
	"github.com/paketo-buildpacks/packit"
)

func main() {
	packit.Build(poetry.Build())
}
