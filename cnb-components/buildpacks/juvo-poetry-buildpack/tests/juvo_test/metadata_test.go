package juvo_test

import (
	"fmt"
	"strings"

	"github.com/juris-futura/juvo-poetry-buildpack/juvo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metadata", func() {
	Context("With a well-formed toml", func() {
		It("Parses the file", func() {
			poetryVer := ""
			pythonVer := ""
			fixture := juvo.BPMetadata{
				PythonVersion: pythonVer,
				PoetryVersion: poetryVer,
			}
			input := strings.NewReader(fmt.Sprintf(`
				[metadata]
				  [[metadata.dependencies]]
				    name = "python"
					version = "%s"
				  [[metadata.dependencies]]
				    name = "poetry"
				    version = "%s"
			`, pythonVer, poetryVer))
			result, err := juvo.ReadMetadata(input)

			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(fixture))
		})
	})
})
