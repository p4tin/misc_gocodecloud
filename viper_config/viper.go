package main

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/viper"
)

type Collector struct {
	Name                          string `mapstructure:"name"`
	CloudType                     string `mapstructure:"type"`
	AwsRegion                     string `mapstructure:"aws_region"`
	OrderFulfillmentQueueURL      string `mapstructure:"order_fulfillment_queue_url"`
	OrderFulfillmentErrorQueueURL string `mapstructure:"order_fulfillment_error_queue_url"`
}

type Configuration struct {
	// Common App variables
	AppName          string `mapstructure:"app_name"`
	GitHash          string `mapstructure:"github_hash"`
	Branch           string `mapstructure:"github_branch"`
	BuildNumber      string `mapstructure:"build_number"`
	Environment      string `mapstructure:"environment"`
	MaxSubscribers   int    `mapstructure:"max_collectors"`
	MaxWorkers       int    `mapstructure:"max_workers"`
	MaxRetries       int    `mapstructure:"max_retries"`
	RequeueDelayTime int64  `mapstructure:"requeue_delay_time"`
	NewRelicLicense  string `mapstructure:"new_relic_license"`

	// Mongo Info
	MongoHost        string `mapstructure:"mongo_host"`
	MongoReadConcern int    `mapstructure:"mongo_read_concern"`

	// AWS SQS Info
	// AwsURL is only set for goaws and local testing, for dev, stage, prod this needs be blank as it is used in the creation of the QueueURLs
	// All the fields here are set for goaws except for AwsURL which is set for dev, stage, prod
	GoAwsEndpoint      string      `mapstructure:"go_aws_endpoint"`
	AwsAccessKeyID     string      `mapstructure:"aws_access_key_id"`
	AwsSecretAccessKey string      `mapstructure:"aws_secret_access_key"`
	MaxPollingPeriod   int64       `mapstructure:"max_polling_period"` // MaxPollingPeriod for AWS failures in millis
	Collectors         []Collector `mapstructure:"collectors"`

	//Tibco - OrderSubmitURL
	OrderSubmitURL     string `mapstructure:"order_submit_url"`
	OrderSubmitTimeout int    `mapstructure:"order_submit_timeout"` // timeout in secs no default

	//Order Number Generator
	OrderNumberGeneratorURL     string `mapstructure:"order_number_generator_url"`
	OrderNumberGeneratorTimeout int    `mapstructure:"order_number_generator_timeout"` // timeout in secs no default
}

func main() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./viper_config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error reading config file, %v", err)
		os.Exit(0)
	}

	//fmt.Printf("%+v\n", viper.GetViper().AllSettings())

	config := new(Configuration)
	if err := viper.Unmarshal(config); err != nil {
		fmt.Printf("error reading config file, %v", err)
		os.Exit(0)
	}

	var cloudType, awsRegion, orderFulfillmentQueueURL, orderFulfillmentErrorQueueURL string

	for i, collector := range config.Collectors {
		if len(collector.Name) == 0 {
			fmt.Printf("name not defined for host %d", i)
			os.Exit(0)
		}

		if cloudType = os.Getenv(strings.ToUpper(collector.Name) + "_" + "TYPE"); len(cloudType) > 0 {
			config.Collectors[i].CloudType = cloudType
		} else if len(collector.CloudType) == 0 {
			fmt.Printf("cloudType not defined for %s", collector.Name)
			os.Exit(0)
		}

		if awsRegion = os.Getenv(strings.ToUpper(collector.Name) + "_" + "AWS_REGION"); len(awsRegion) > 0 {
			config.Collectors[i].AwsRegion = awsRegion
		} else if len(collector.AwsRegion) == 0 {
			fmt.Printf("awsRegion not defined for %s", collector.Name)
			os.Exit(0)
		}

		queue_url := strings.ToUpper(collector.Name) + "_" + "ORDER_FULFILLMENT_QUEUE_URL"
		if orderFulfillmentQueueURL = os.Getenv(queue_url); len(orderFulfillmentQueueURL) > 0 {
			config.Collectors[i].OrderFulfillmentQueueURL = orderFulfillmentQueueURL
		} else if len(collector.OrderFulfillmentQueueURL) == 0 {
			fmt.Printf("orderFulfillmentQueueURL not defined for %s", collector.Name)
			os.Exit(0)
		}

		if orderFulfillmentErrorQueueURL = os.Getenv(strings.ToUpper(collector.Name) + "_" + "ORDER_FULFILLMENT_ERROR_QUEUE_URL"); len(orderFulfillmentErrorQueueURL) > 0 {
			config.Collectors[i].OrderFulfillmentErrorQueueURL = orderFulfillmentErrorQueueURL
		} else if len(collector.OrderFulfillmentErrorQueueURL) == 0 {
			fmt.Printf("pasorderFulfillmentErrorQueueURLsword not defined for %s", collector.Name)
			os.Exit(0)
		}
	}

	fmt.Printf("%+v\n", config)
}
