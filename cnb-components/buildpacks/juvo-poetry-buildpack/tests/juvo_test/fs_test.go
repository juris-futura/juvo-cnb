package juvo_test

import (
	"os"
	"path/filepath"

	"github.com/juris-futura/juvo-poetry-buildpack/juvo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fs", func() {
	var fname string
	var tmpdir string

	BeforeEach(func() {
		var err error
		tmpdir, err = os.MkdirTemp("", "fs_test")
		Expect(err).ToNot(HaveOccurred())

		fname = filepath.Join(tmpdir, "foo")
		f, err := os.Create(fname)
		Expect(err).ToNot(HaveOccurred())
		f.Close()
	})

	AfterEach(func() {
		Expect(os.RemoveAll(tmpdir)).To(Succeed())
	})

	Context("When file exists", func() {
		It("Returns true", func() {
			fs := juvo.PhysicalFs{}
			Expect(fs.FileExists(fname)).To(BeTrue())
		})
	})

	Context("When file does not exist", func() {
		It("Returns false", func() {
			fs := juvo.PhysicalFs{}
			invalidname := filepath.Join(tmpdir, "invalid")
			Expect(fs.FileExists(invalidname)).To(BeFalse())
		})
	})
})
