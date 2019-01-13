package main

import (
//	"bytes"
	"fmt"
	"log"
	"time"
	"os/exec"
//	"strings"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"os/exec"
)

type Config struct {
	MaxBots int `yaml:"maxbots"`
	Interval int `yaml:""interval`
	Bots []string `yaml:"bots"`
}


func main() {
	dat, err := ioutil.ReadFile("yaml/go2bot_runner.yaml")
	check(err)

	config := Config{}
	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	numbots := len(config.Bots)
	fmt.Println(config.MaxBots)
	fmt.Println(config.Interval)
	fmt.Printf("Numbots: %d\n", numbots)
	for _, bot := range config.Bots {
		fmt.Println("\t - " + bot + "\\AvdGo2Bot")
	}

	bot_index := 0
	for {
		running_bot:= make([]exec.Cmd, 0, 0)
		for x := 0; x < config.MaxBots; x++ {
			cmd := exec.Command(fmt.Sprintf("%s%s", config.Bots[bot_index], "\\AvdGo2Bot")
			cmd.Dir = config.Bots[bot_index]
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Running bot %d with command %s%s\n", bot_index, config.Bots[bot_index], "\\AvdGo2Bot")
			running_bot = append(running_bot, cmd)
			bot_index++
			if bot_index >= numbots {
				bot_index = 0
			}
		}
		fmt.Printf("Waiting Interval %d\n", config.Interval)
		time.Sleep(5 * time.Second)
		for _, cmd := range running_bot {
			fmt.Printf("Killing bot %s\n", cmd.Path)
			if err := cmd.Process.Kill(); err != nil {
				log.Fatal("failed to kill: ", err)
			}
		}
	}

	//cmd := exec.Command("tr", "a-z", "A-Z")
	//cmd.Stdin = strings.NewReader("bla bla bla")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err = cmd.Run()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("in all caps: %q\n", out.String())
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}
