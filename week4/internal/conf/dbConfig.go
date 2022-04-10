package conf

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type MysqlConfig struct {
	SourceType int `yaml:"data.database.driver"`
	SourceInfo int `yaml:"data.database.source"`
}

type RedisConfig struct {
	Addr         int `yaml:"data.redis.addr"`
	ReadTimeOut  int `yaml:"data.redis.read_timeout"`
	WriteTimeOut int `yaml:"data.redis.write_timeout"`
}

func (c *MysqlConfig) LoadFile(path string) error {
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

func (c *RedisConfig) LoadFile(path string) error {
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
