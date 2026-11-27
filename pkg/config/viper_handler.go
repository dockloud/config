package config

import (
	"github.com/spf13/viper"
)

type ViperHandler[T any] struct {
	path     string
	filename string
	format   string
}

func NewViperHandler[T any](path, filename, format string) *ViperHandler[T] {
	return &ViperHandler[T]{
		path:     path,
		filename: filename,
		format:   format,
	}
}

func (h *ViperHandler[T]) GetConfig() (*T, error) {
	v := viper.New()

	v.SetConfigName(h.filename)
	v.SetConfigType(h.format)
	v.AddConfigPath(h.path)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg T
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
