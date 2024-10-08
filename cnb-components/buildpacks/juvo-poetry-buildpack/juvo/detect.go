package juvo

import (
	"fmt"
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

type Metadata struct {
	Build   bool   `toml:"build"`
	Launch  bool   `toml:"launch"`
	Version string `toml:"version"`
}

func Detect(fs Fs) packit.DetectFunc {
	shouldDetect := func(workdir string) error {
		path := filepath.Join(workdir, "pyproject.toml")
		exists, err := fs.FileExists(path)
		if err != nil {
			return err
		}

		if exists {
			return nil
		}

		err = fmt.Errorf("shouldDetect: File not found %s", path)
		return err
	}

	return func(ctx packit.DetectContext) (packit.DetectResult, error) {
		if err := shouldDetect(ctx.WorkingDir); err != nil {
			return packit.DetectResult{}, err
		}

		fmt.Println("Fetching Version Info . . .")
		versions, err := fs.ParseMetadataFromFile(ctx.CNBPath)
		if err != nil {
			return packit.DetectResult{}, err
		}
		fmt.Printf("Poetry Version: %s\n", versions.PoetryVersion)
		fmt.Printf("Python Version: %s\n", versions.PythonVersion)

		requires := []packit.BuildPlanRequirement{
			{
				Name: "cpython",
				Metadata: Metadata{
					Build:  true,
					Launch: true,
				},
			},
			{
				Name: "poetry",
				Metadata: Metadata{
					Version: versions.PoetryVersion, // use the verson configured in buildpack.toml
					Build:   true,
					Launch:  true,
				},
			},
			{
				Name: "juvo",
			},
		}

		provides := []packit.BuildPlanProvision{
			{
				Name: "juvo",
			},
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: provides,
				Requires: requires,
			},
		}, nil
	}
}
