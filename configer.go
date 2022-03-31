package _config

import (
	"fmt"
	"github.com/gozelle/_toml"
)

type Configer struct {
	content []byte
}

func (p Configer) Bind(target interface{}) (err error) {
	err = _toml.Unmarshal(p.content, target)
	if err != nil {
		err = fmt.Errorf("bind config error: %s", err)
		return
	}
	return
}
