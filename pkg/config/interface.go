package config

type Handler[T any] interface {
	GetConfig() (*T, error)
}
