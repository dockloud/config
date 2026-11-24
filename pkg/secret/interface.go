package secret

type Handler interface {
	ReadSecret(name string) (string, error)
}
