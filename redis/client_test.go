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

		Describe(".lookupAlias", func() {
			var redisClient *client

			BeforeEach(func() {
				redisClient = newClient(nil)
			})

			Context("when alias exists", func() {
				BeforeEach(func() {
					CommandAlias("cmd", "alias")(redisClient)
				})

				It("returns the alias", func() {
					Expect(redisClient.lookupAlias("CMD")).To(Equal("alias"))
				})

				It("is not case-sensitive", func() {
					Expect(redisClient.lookupAlias("cMd")).To(Equal("alias"))
				})
			})

			Context("when the alias does not exists", func() {
				It("returns the command itself", func() {
					Expect(redisClient.lookupAlias("cmd")).To(Equal("cmd"))
				})
			})
		})
	})

})
