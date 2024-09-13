package juvo

import (
	"fmt"

	"github.com/paketo-buildpacks/packit"
)

func Build() packit.BuildFunc {
	return func(ctx packit.BuildContext) (packit.BuildResult, error) {
		juvoLayer, err := ctx.Layers.Get("juvo")
		if err != nil {
			return packit.BuildResult{}, err
		}

		juvoLayer, err = juvoLayer.Reset()
		if err != nil {
			return packit.BuildResult{}, err
		}

		juvoLayer.Launch = true

		fmt.Println("Installing Virtual Env . . .")

		poetryConfig := CommandDescriptor{
			Cmd:  "poetry",
			Args: []string{"config", "virtualenvs.in-project", "true"},
		}

		if err = ExecuteStep(poetryConfig); err != nil {
			return packit.BuildResult{}, err
		}

		var poetryInstall = PoetryInstall{
			OnlyMain:       true,
			KeyFilePath:    "/platform/bindings/git/id_rsa",
			FallbackEnvVar: "PRIV_SSH_KEY",
		}

		if err = ExecuteStep(poetryInstall); err != nil {
			return packit.BuildResult{}, err
		}

		return packit.BuildResult{
			Layers: []packit.Layer{juvoLayer},
		}, nil
	}
}
