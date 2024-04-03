package poetry

import (
	"os"
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

func Detect() packit.DetectFunc {
	return func(ctx packit.DetectContext) (packit.DetectResult, error) {
		if err := shouldDetect(ctx.WorkingDir); err != nil {
			return packit.DetectResult{}, err
		}
		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{{Name: "poetry"}},
				Requires: []packit.BuildPlanRequirement{{Name: "poetry"}},
			},
		}, nil
	}
}

func shouldDetect(workdir string) error {
	path := filepath.Join(workdir, "pyproject.toml")
	_, err := os.Stat(path)
	return err
}
