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
			redisClient    Client
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
			redisClient, connectErr = Connect("foobar", 1234, options...)
		})

		It("returns a client", func() {
			Expect(redisClient).ToNot(BeNil())
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

				It("closes the connection", func() {
					Expect(conn.ReceivedCloseCalls()).To(Equal(1))
				})
			})
		})

		Context("with generic options", func() {
			var cnt int

			BeforeEach(func() {
				inc := func(c *client) {
					cnt++
				}

				options = []option{inc, inc, inc}
			})

			It("calls all options", func() {
				Expect(cnt).To(Equal(3))
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

	Describe("Options", func() {
		var redisClient *client

		BeforeEach(func() {
			redisClient = newClient(nil)
		})

		Describe("Password", func() {
			It("sets the password on the client", func() {
				Password("some-password")(redisClient)
				Expect(redisClient.auth).To(Equal("some-password"))
			})
		})

		Describe("CommandAlias", func() {
			It("adds the alias with an uppercase command", func() {
				CommandAlias("cmd", "alias")(redisClient)

				Expect(redisClient.aliases).To(
					Equal(map[string]string{"CMD": "alias"}),
				)
			})

			It("can be called multiple times to add more than one alias", func() {
				CommandAlias("cmd1", "alias1")(redisClient)
				CommandAlias("cmd2", "alias2")(redisClient)

				Expect(redisClient.aliases).To(
					Equal(map[string]string{
						"CMD1": "alias1",
						"CMD2": "alias2",
					}),
				)
			})
		})

	})
})
