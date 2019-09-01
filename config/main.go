package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

// S3Config is a struct containing the S3 properties
type S3Config struct {
	Bucket string `json:"bucket"`
	Region string `json:"region"`
}

// AWSConfig is a struct containing the credentials of an IAM user
type AWSConfig struct {
	Key    string   `json:"key"`
	Secret string   `json:"secret"`
	S3     S3Config `json:"s3"`
}

// Config is a configuration struct
type Config struct {
	AWS AWSConfig `json:"aws"`
}

// GetConfig returns a configuration object
func GetConfig() Config {
	home, _ := homedir.Dir()
	joercPath := path.Join(home, ".joerc.json")

	if _, err := os.Stat(joercPath); os.IsNotExist(err) {
		fmt.Println("Please create a .joerc.json file in your home directory to use this command.")
		os.Exit(1)
	}

	joerc, err := os.Open(joercPath)

	if err != nil {
		panic("Failed to open ~/.joerc.json")
	}

	defer joerc.Close()

	contents, _ := ioutil.ReadAll(joerc)

	var config Config
	json.Unmarshal(contents, &config)

	if err != nil {
		panic("Failed to read configuration file.")
	}

	return config
}
