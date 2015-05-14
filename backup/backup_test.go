package backup_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/st3v/redis-tools/backup"
)

var _ = Describe("backup", func() {

	Describe("Backup", func() {

		Describe(".Run", func() {
			var subject *backup.Backup

			BeforeEach(func() {
				subject = backup.New()
			})

			It("does not return an error", func() {
				err := subject.Run()
				Expect(err).ToNot(HaveOccurred())
			})

			/**

			It("connects to redis", func() {

			})

			It("performs a BGSAVE", func() {

			})

			It("copies the dump", func() {

			})

			It("names the dump accordingly", func() {

			})

			*/

		})

	})

})
