package redis

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/st3v/redis-tools/redis/fake"

	redigo "github.com/garyburd/redigo/redis"
)

var _ = Describe("redis", func() {
	Describe(".Connect", func() {
		var (
			dialedProtocol string
			dialedAddress  string
			connectErr     error
			client         Client
			conn           *fake.Conn
			options        []option
		)

		BeforeEach(func() {
			conn = &fake.Conn{}

			dialedProtocol = ""
			dialedAddress = ""

			dialer = func(protocol, address string) (redigo.Conn, error) {
				dialedProtocol = protocol
				dialedAddress = address
				return conn, nil
			}

			options = []option{}
		})

		JustBeforeEach(func() {
			client, connectErr = Connect("foobar", 1234, options...)
		})

		It("returns a client", func() {
			Expect(client).ToNot(BeNil())
		})

		It("does not return an error", func() {
			Expect(connectErr).ToNot(HaveOccurred())
		})

		It("opens a connection with the correct protocol", func() {
			Expect(dialedProtocol).To(Equal("tcp"))
		})

		It("opens a connection with the correct address", func() {
			Expect(dialedAddress).To(Equal("foobar:1234"))
		})

		Context("without password option", func() {
			It("does not authenticate with redis", func() {
				doCalls := conn.ReceivedDoCalls()
				Expect(doCalls).To(BeEmpty())
			})
		})

		Context("with password option", func() {
			BeforeEach(func() {
				options = []option{
					Password("some-password"),
				}
			})

			It("authenticates with redis", func() {
				doCalls := conn.ReceivedDoCalls()
				authArgs, authenticated := doCalls["AUTH"]
				Expect(doCalls).To(HaveLen(1))
				Expect(authenticated).To(BeTrue())
				Expect(authArgs).To(Equal([]interface{}{"some-password"}))
			})

			Context("when an authentication error occurs", func() {
				BeforeEach(func() {
					conn.ExpectedDoErr = errors.New("auth-error")
				})

				It("returns the error", func() {
					Expect(connectErr).To(HaveOccurred())
					Expect(connectErr.Error()).To(ContainSubstring("auth-error"))
				})
			})
		})

		Context("when a connection error occurs", func() {
			BeforeEach(func() {
				dialer = func(string, string) (redigo.Conn, error) {
					return nil, errors.New("dial-error")
				}
			})

			It("returns the error", func() {
				Expect(connectErr).To(HaveOccurred())
				Expect(connectErr.Error()).To(ContainSubstring("dial-error"))
			})
		})
	})
})
