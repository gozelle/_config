package _config

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"github.com/gozelle/_toml"
	"io/ioutil"
	"os"
)

func Load(filepath string, ptr interface{}) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		err = fmt.Errorf("load config open file error: %s", err)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("load config read file error: %s", err)
		return
	}

	err = _toml.Unmarshal(bytes, ptr)
	if err != nil {
		err = fmt.Errorf("load config bind config error: %s", err)
		return
	}

	return
}

func Dump(ptr interface{}) {
	dump.P(ptr)
}
