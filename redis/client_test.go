package redis

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/st3v/redis-tools/redis/fake"
)

var _ = Describe("redis", func() {
	Describe("Client", func() {
		var (
			conn    *fake.Conn
			subject Client
		)

		BeforeEach(func() {
			conn = &fake.Conn{}
			subject = &client{
				conn: conn,
			}
		})

		Describe(".Close", func() {
			var closeErr error

			JustBeforeEach(func() {
				closeErr = subject.Close()
			})

			It("does not return an error", func() {
				Expect(closeErr).ToNot(HaveOccurred())
			})

			It("closes the redis connection", func() {
				Expect(conn.ReceivedCloseCalls()).To(Equal(1))
			})

			Context("when closing the connection fails", func() {
				BeforeEach(func() {
					conn.ExpectedCloseErr = errors.New("some-error")
				})

				It("returns an error", func() {
					Expect(closeErr).To(MatchError("some-error"))
				})
			})
		})
	})

})
