package poetry

import (
	"fmt"
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

func Build() packit.BuildFunc {
	return func(ctx packit.BuildContext) (packit.BuildResult, error) {
		// Read the content of buildpack.toml. Well find poetry dep there
		var input = MetaInput{
			BuildpackMetadataPath: filepath.Join(ctx.CNBPath, "buildpack.toml"),
		}
		fmt.Println("Fetching Poetry Version . . .")
		poetryVersion, err := input.ReadMetadata()
		if err != nil {
			return packit.BuildResult{}, err
		}
		fmt.Printf("Poetry Version: %s\n", poetryVersion)

		poetryLayer, err := ctx.Layers.Get("poetry")
		if err != nil {
			return packit.BuildResult{}, err
		}

		poetryLayer, err = poetryLayer.Reset()
		if err != nil {
			return packit.BuildResult{}, err
		}

		poetryLayer.Launch = true

		fmt.Println("Installing Virtual Env . . .")
		var poetryInstall = PoetryInstall{
			OnlyMain:    true,
			KeyFilePath: "/platform/bindings/git/id_rsa",
		}
		if err = ExecuteStep(poetryInstall); err != nil {
			return packit.BuildResult{}, err
		}

		return packit.BuildResult{
			Layers: []packit.Layer{poetryLayer},
		}, nil
	}
}
