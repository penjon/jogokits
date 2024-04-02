package env

import (
	"fmt"
	"testing"
)

type ConfigChild struct {
	Tag string `env:"tag" env_default:"tag_default"`
}
type Config struct {
	Name  string       `env:"mongodb"`
	Desc  string       `env:"desc" env_default:""`
	Num   int          `env_default:"12"`
	Size  uint         `env_default:"24"`
	Child *ConfigChild `env_prefix:"test"`
}

func TestLoad(t *testing.T) {
	c := &Config{}

	Load(c, "")

	fmt.Printf("%v\n", c)
}
