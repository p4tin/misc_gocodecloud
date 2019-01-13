package main
import (
	"os/user"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Username string
	DefaultNotebook string
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	fmt.Printf("%+v\n", usr )

	if _, err := os.Stat(usr.HomeDir+"/.notabene"); os.IsNotExist(err) {
		os.Mkdir(usr.HomeDir+"/.notabene", 0777)
		fmt.Println("Creating direcory for .notabene")
	} else {
		fmt.Println(".notabene exits.")
	}

	if _, err := os.Stat(usr.HomeDir+"/.notabene/config"); os.IsNotExist(err) {
		c := Config{Username: usr.Username, DefaultNotebook: "Home"}
		d, err := yaml.Marshal(&c)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- c dump:\n%s\n\n", string(d))

		ioutil.WriteFile(usr.HomeDir+"/.notabene/config", d, 0644)
		fmt.Println("Creating config file")
	} else {
		fmt.Println("config exists")
	}
}
