package backup

const (
	hostname = "localhost"
	port     = 6379
)

type Backup struct {
	port int
	host string
	auth string
}

func (b *Backup) Run() error {
	return nil
}

func New(options ...option) *Backup {
	b := &Backup{
		port: port,
		host: hostname,
	}

	for _, option := range options {
		option(b)
	}

	return b
}
