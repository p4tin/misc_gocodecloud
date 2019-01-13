package main

import (
	"fmt"
	"github.com/apex/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Name            []string `yaml:"name"`
	ListenPort      []string `yaml:"listen_port"`
	DefaultLanguage []string `yaml:"default_language"`
	DefaultChannel  []string `yaml:"default_channel"`
}

func main() {
	dat, err := ioutil.ReadFile("yaml/config.yaml")
	check(err)

	config := Config{}
	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(GetEnvOrDefault(config.Name[0], config.Name[1]))
	fmt.Println(GetEnvOrDefault(config.ListenPort[0], config.ListenPort[1]))
	fmt.Println(GetEnvOrDefault(config.DefaultLanguage[0], config.DefaultLanguage[1]))
	fmt.Println(GetEnvOrDefault(config.DefaultChannel[0], config.DefaultChannel[1]))
}

func GetEnvOrDefault(name string, def string) string {
	val := os.Getenv(name)
	if val != "" {
		return val
	}
	return def
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
