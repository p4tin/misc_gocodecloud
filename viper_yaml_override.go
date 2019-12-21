package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/davecgh/go-spew/spew"
)

type Settings struct {
	AppName         string `yaml:"app_name"`
	Environment     string `yaml:"environment"`
	MaxPollers      string `yaml:"max_pollers"`
	MaxReceivers    string `yaml:"max_receivers"`
	PollingPeriod   string `yaml:"polling_period"`
	NewrelicLicense string `yaml:"newrelic_license"`
	Mongo           struct {
		Host                       string `yaml:"host"`
		UseSsl                     string `yaml:"use_ssl"`
		CertPath                   string `yaml:"cert_path"`
		Rootca                     string `yaml:"rootca"`
		ClientKey                  string `yaml:"client_key"`
		ClientCertificate          string `yaml:"client_certificate"`
		Database                   string `yaml:"database"`
		InventoryCollection        string `yaml:"inventory_collection"`
		PooldefinitionCollection   string `yaml:"pooldefinition_collection"`
		SkuDetailCollection        string `yaml:"sku_detail_collection"`
		ProductfilterCollection    string `yaml:"productfilter_collection"`
		OrderSkuQuantityCollection string `yaml:"order_sku_quantity_collection"`
	} `yaml:"mongo"`
	Aws struct {
		URL             string `yaml:"url"`
		ClientID        string `yaml:"client_id"`
		RegionID        string `yaml:"region_id"`
		AccessKeyID     string `yaml:"access_key_id"`
		SecretAccessKey string `yaml:"secret_access_key"`
		InventoryQueue  string `yaml:"inventory_queue"`
	} `yaml:"aws"`
}

var ProcessedYamlConfig string

func main() {
	file, err := os.Open("viper_yaml.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newValue := getEnvOrDefault(scanner.Text())
		ProcessedYamlConfig = fmt.Sprintf("%s\n%s", ProcessedYamlConfig, newValue)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	spew.Dump(ProcessedYamlConfig)

	settings := new(Settings)

	err = yaml.Unmarshal([]byte(ProcessedYamlConfig), settings)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", settings)
}

func getEnvOrDefault(line string) string {
	parts := strings.Split(line, ":")
	if len(parts) < 2 || parts[1] == "" {
		return line
	}
	values := strings.Split(strings.TrimSpace(parts[1]), ",")
	value := strings.TrimSpace(values[0])
	if strings.HasPrefix(value, "<%= ENV['") && strings.HasSuffix(value, "'] %>") {
		env := strings.TrimSuffix(strings.TrimPrefix(value, "<%= ENV['"), "'] %>")
		value = os.Getenv(env)
		if value == "" {
			value = strings.TrimSpace(values[1])
		}
	} else {
		value = strings.TrimSpace(values[1])
	}
	index := strings.Index(line, "<%= ENV['")

	newline := strings.Replace(line,
		line[index:len(line)], value, 1)
	return newline
}
