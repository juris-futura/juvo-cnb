package juvo_test

import (
	"github.com/juris-futura/juvo-poetry-buildpack/juvo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/packit"
)

var _ = Describe("Build", func() {
	It("Returns a Juvo Layer", func() {
		result, err := juvo.Build(MockExecutor{})(packit.BuildContext{})
		Expect(err).ToNot(HaveOccurred())

		var layers []string
		for _, l := range result.Layers {
			layers = append(layers, l.Name)
		}
		Expect(layers).To(ContainElement(Equal("juvo")))
	})
})
