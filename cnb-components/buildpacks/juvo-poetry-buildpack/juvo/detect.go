package juvo

import (
	"fmt"
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

func Detect(fs Fs) packit.DetectFunc {
	shouldDetect := func(workdir string) error {
		path := filepath.Join(workdir, "pyproject.toml")
		if fs.FileExists(path) {
			return nil
		}
		err := fmt.Errorf("shouldDetect: File not found %s", path)
		return err
	}

	return func(ctx packit.DetectContext) (packit.DetectResult, error) {
		if err := shouldDetect(ctx.WorkingDir); err != nil {
			return packit.DetectResult{}, err
		}

		requires := []packit.BuildPlanRequirement{
			{
				Name: "poetry",
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
