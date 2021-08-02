package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type StorageType string

const (
	GCS StorageType = "GCS"
	S3 = "S3"
)

type StorageConfiguration interface {
	LoadConfiguration()
	GetType() StorageType
}

type GCPSecretConfiguration struct {
	Id   string `yaml:"id,omitempty"`
	Name string `yaml:"name"`
}

type GCSConfiguration struct {
	BucketName string                 `yaml:"bucket-name"`
	Secret     GCPSecretConfiguration `yaml:"secret"`
}

type GCPConfiguration struct {
	GCS GCSConfiguration `yaml:"gcs"`
}

func (c *GCPConfiguration) GetType() StorageType {
	return GCS
}

func (c *GCPConfiguration) LoadConfiguration() {
	err := loadConfiguration("configs/gcp.yaml", c)
	if err != nil {
		panic(err)
	}
}

func loadConfiguration(filename string, out interface{}) error {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("reading yaml file [%s] returned an error. got: %v", filename, err)
	}
	// expand environment variables
	yamlFile = []byte(os.ExpandEnv(string(yamlFile)))
	err = yaml.Unmarshal(yamlFile, out)
	if err != nil {
		return fmt.Errorf("unmarshalling yaml bytes failed. got: %v", err)
	}
	return nil
}
