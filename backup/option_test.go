package backup

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("backup.New", func() {
	var backup *Backup

	Context("without options", func() {
		BeforeEach(func() {
			backup = New()
		})

		It("sets the default redis port 6379", func() {
			Expect(backup.port).To(Equal(6379))
		})
	})

	Context("with options", func() {
		BeforeEach(func() {
			backup = New(
				Host("foobar"),
				Port(1234),
				Auth("password"),
			)
		})

		It("sets the correct port", func() {
			Expect(backup.port).To(Equal(1234))
		})
	})
})
