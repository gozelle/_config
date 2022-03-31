package _config

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Load(filepath string, config interface{}) (configer *Configer, err error) {
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
	
	configer = new(Configer)
	configer.content = bytes
	return
}
