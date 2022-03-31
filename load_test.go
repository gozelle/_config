package _config

import (
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

type Config struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func TestLoad(t *testing.T) {
	pwd, err := os.Getwd()
	require.NoError(t, err)
	config := Config{}
	err = Load(filepath.Join(pwd, "/test.toml"), &config)
	require.NoError(t, err)
	require.Equal(t, "127.0.0.1", config.Host)
	require.Equal(t, 22, config.Port)
	dump.P(config)
	dump.V(config)
	dump.Println(config)
}
