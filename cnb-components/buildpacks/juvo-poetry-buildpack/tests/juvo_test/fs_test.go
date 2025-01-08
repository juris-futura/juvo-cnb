package juvo_test

import (
	"os"
	"path/filepath"

	"github.com/juris-futura/juvo-poetry-buildpack/juvo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fs", func() {
	Describe("ParseMetadataFromFile", func() {
		var fname string
		var tmpdir string

		BeforeEach(func() {
			var err error
			tmpdir, err = os.MkdirTemp("", "fs_test")
			Expect(err).ToNot(HaveOccurred())

			fname = filepath.Join(tmpdir, "buildpack.toml")
			f, err := os.Create(fname)
			defer f.Close()
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			Expect(os.RemoveAll(tmpdir)).To(Succeed())
		})

		Context("On a well-formed toml", func() {
			It("Parses the file", func() {
				fs := juvo.PhysicalFs{}
				_, err := fs.ParseMetadataFromFile(tmpdir)

				Expect(err).ToNot(HaveOccurred())
			})
		})

	})
	Describe("FileExists", func() {
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
				exists, err := fs.FileExists(fname)
				Expect(err).ToNot(HaveOccurred())
				Expect(exists).To(BeTrue())
			})
		})

		Context("When file does not exist", func() {
			It("Returns false", func() {
				fs := juvo.PhysicalFs{}
				invalidname := filepath.Join(tmpdir, "invalid")
				exists, err := fs.FileExists(invalidname)
				Expect(err).ToNot(HaveOccurred())
				Expect(exists).To(BeFalse())
			})
		})
	})
})
