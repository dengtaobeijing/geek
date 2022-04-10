package conf

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type ServerConfig struct {
	Server0Port int `yaml:"appPort1"`
	Server1Port int `yaml:"appPort2"`
}

//server:
//  http:
//    appPort1: 7070
//    appPort2: 7071

func (c *ServerConfig) LoadFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, c)
	if err != nil {

		return err
	}
	return nil
}
