package main

import (
	"fmt"

	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	return v, err
}

func main() {
	v1, err := readConfig("config", map[string]interface{}{
		"port":     9090,
		"hostname": "localhost",
		"auth": map[string]string{
			"username": "titpetric",
			"password": "12fa",
		},
	})
	if err != nil {
		panic(fmt.Errorf("Error when reading config: %v\n", err))
	}

	v1.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	port := v1.GetInt("port")
	hostname := v1.GetString("hostname")
	auth := v1.GetString("auth.password")
	username := v1.GetString("auth.username")

	fmt.Printf("Reading config for port = %d\n", port)
	fmt.Printf("Reading config for hostname = %s\n", hostname)
	fmt.Printf("Reading config for auth = %s\n", auth)
	fmt.Printf("Reading config for auth.username = %s\n", username)

	spew.Dump(v1)
}
