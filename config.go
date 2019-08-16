package zconfig

import (
	"errors"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var fileNotFound = errors.New("error reading yaml file")

func Load(t interface{}, ymls ...string) error {
	for _, v := range ymls {
		err := readFile(t, v)
		if err == fileNotFound {
			log.Println("file reading skipped")
		}
	}
	return readEnv(t)
}

func readFile(cfg interface{}, ymlPath string) error {
	f, err := os.Open(ymlPath)
	if err != nil {
		return fileNotFound
	}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return fmt.Errorf("error decoding %s: %v", ymlPath, err)
	}
	return nil
}

func readEnv(cfg interface{}) error {
	return envconfig.Process("", cfg)
}
