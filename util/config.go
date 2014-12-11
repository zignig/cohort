package util

import (
	"bytes"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Ref    string `toml:"ref"`
	Path   string `toml:"path"`
	Banner string `toml:"banner"`
}

func GetConfig(path string) (c *Config) {
	if _, err := toml.DecodeFile(path, &c); err != nil {
		fmt.Println("this is an error")
		fmt.Println(err)
		return
	}
	return
}

func (c *Config) SaveConfig() (err error) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String)
	return err
}
