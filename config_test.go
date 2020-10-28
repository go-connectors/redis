package redis

import (
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	var cases = []struct {
		name   string
		config Config
		err    error
	}{
		{name: "valid", config: Config{Addr: "127.0.0.1:6379"}, err: nil},
		{name: "empty addr", config: Config{}, err: ErrEmptyAddr},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if err != tt.err {
				t.Errorf("got %q, want %q", err, tt.err)
			}
		})
	}
}
