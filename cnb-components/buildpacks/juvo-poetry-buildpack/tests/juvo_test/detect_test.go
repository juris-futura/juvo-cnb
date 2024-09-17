package juvo_test

import (
	poetry "github.com/juris-futura/juvo-poetry-buildpack/juvo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/packit"
)

var _ = Describe("Detect", func() {
	var validFs *MockFs
	var invalidFs *MockFs
	var context packit.DetectContext

	BeforeEach(func() {
		validFs = &MockFs{
			Files: map[string]string{"pyproject.toml": ""},
		}
		invalidFs = &MockFs{
			Files: map[string]string{},
		}
	})

	Context("When pyproject.toml exists", func() {
		It("passes detection", func() {
			result, err := poetry.Detect(validFs)(context)
			Expect(err).NotTo(HaveOccurred())

			var requirements []string
			for _, req := range result.Plan.Requires {
				requirements = append(requirements, req.Name)
			}

			var providers []string
			for _, prov := range result.Plan.Provides {
				providers = append(providers, prov.Name)
			}

			Expect(requirements).To(HaveExactElements("cpython", "poetry", "juvo"))
			Expect(providers).To(HaveExactElements("juvo"))
		})
	})

	Context("When pyproject.toml does not exists", func() {
		It("fails detection", func() {
			_, err := poetry.Detect(invalidFs)(context)
			Expect(err).To(HaveOccurred())
		})
	})
})
