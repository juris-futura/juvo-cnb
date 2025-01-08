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

		requires := []packit.BuildPlanRequirement{
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
