package juvo_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/juris-futura/juvo-poetry-buildpack/juvo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PoetryInstall", func() {
	var poetryInstall juvo.PoetryInstall
	var tmpdir string
	var keyfile string

	BeforeEach(func() {
		var err error
		tmpdir, err = os.MkdirTemp("", "fs_test")
		Expect(err).ToNot(HaveOccurred())

		keyfile = filepath.Join(tmpdir, "keyfile")
		poetryInstall = juvo.PoetryInstall{
			KeyFilePath:    keyfile,
			CtxPath:        tmpdir,
			OnlyMain:       false,
			FallbackEnvVar: "ENVVAR",
		}
	})

	AfterEach(func() {
		Expect(os.RemoveAll(tmpdir)).To(Succeed())
	})

	Describe("Command Descriptor", func() {
		var fixture juvo.CommandDescriptor

		BeforeEach(func() {
			sshcmd := fmt.Sprintf("ssh -i %s/privkey -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -F /dev/null", tmpdir)
			fixture = juvo.CommandDescriptor{
				Cmd:  "poetry",
				Args: []string{"install", "--no-root"},
				Env: map[string]string{
					"GIT_SSH_COMMAND": sshcmd,
				},
			}

			os.Setenv("ENVVAR", "some_key")
		})

		AfterEach(func() {
			os.Setenv("ENVVAR", "")
		})

		It("returns a command descriptor", func() {
			result, err := poetryInstall.MkCmd()
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(fixture))
		})

		Context("when onlymain is set", func() {
			It("returns a command descriptor", func() {
				poetryInstall.OnlyMain = true
				fixture.Args = append(fixture.Args, "--only=main")

				result, err := poetryInstall.MkCmd()
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal(fixture))
			})
		})
	})

	Describe("Private Key Setting", func() {
		Context("When KEY is not set", func() {
			It("Returns an error", func() {
				_, err := poetryInstall.MkCmd()
				errmsg := fmt.Sprintf("%s not found and ENVVAR not set", keyfile)
				Expect(err).To(MatchError(errmsg))
			})
		})

		Context("When KEY is set", func() {
			key := "12345"
			checkFile := func(fname string) {
				f, err := os.Open(fname)
				defer f.Close()
				Expect(err).ToNot(HaveOccurred())

				b := make([]byte, 5)
				_, err = f.Read(b)
				Expect(err).ToNot(HaveOccurred())

				Expect(string(b)).To(Equal(key))
			}

			Context("When the KEY is set in a file", func() {
				BeforeEach(func() {
					f, err := os.Create(keyfile)
					defer f.Close()
					Expect(err).ToNot(HaveOccurred())
					f.WriteString(key)
				})

				It("set the Key in the file", func() {
					_, err := poetryInstall.MkCmd()
					Expect(err).ToNot(HaveOccurred())
					checkFile(keyfile)
				})
			})

			Context("When the KEY is set in an env var", func() {
				BeforeEach(func() {
					os.Setenv("ENVVAR", key)
				})

				AfterEach(func() {
					os.Setenv("ENVVAR", "")
				})

				It("set the Key in the file", func() {
					_, err := poetryInstall.MkCmd()
					Expect(err).ToNot(HaveOccurred())
					checkFile(filepath.Join(tmpdir, "privkey"))
				})
			})
		})
	})
})
