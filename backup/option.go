package backup

type option func(*Backup)

func Port(port int) option {
	return func(b *Backup) {
		b.port = port
	}
}

func Host(host string) option {
	return func(b *Backup) {
		b.host = host
	}
}

func Auth(password string) option {
	return func(b *Backup) {
		b.auth = password
	}
}
