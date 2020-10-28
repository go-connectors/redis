package redis

import (
	"errors"
	"fmt"
	"time"
)

// Config validation errors.
var (
	// ErrConfigValidation is general config validation error message.
	ErrConfigValidation = errors.New("redis config validation error")

	ErrEmptyAddr = fmt.Errorf("%w: addr is empty", ErrConfigValidation)
)

// Config contains configuration Redis data.
type Config struct {
	Addr         string        `yaml:"addr"`
	Password     string        `yaml:"password"`
	MaxRetries   int           `yaml:"max_retries"`
	DialTimeout  time.Duration `yaml:"dial_timeout"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	PoolSize     int           `yaml:"pool_size"`
	DB           int           `yaml:"db"`
}

// Validate checks required fields.
func (cfg *Config) Validate() error {
	if cfg.Addr == "" {
		return ErrEmptyAddr
	}

	return nil
}
