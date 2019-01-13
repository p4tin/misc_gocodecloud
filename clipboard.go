package main

import (
	"crypto/rand"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"

	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/codegangsta/cli"
)

var key []byte

type User struct {
	Email string
	Pass  string
}

var Sites map[string]map[string]User

func init() {
	content, err := ioutil.ReadFile("./aes.key")
	check(err)
	key, err = base64.StdEncoding.DecodeString(string(content))
	check(err)

	if _, err := os.Stat("data/clipboard.json"); os.IsNotExist(err) {
		users := make(map[string]User, 0)
		Sites = make(map[string]map[string]User, 0)

		encryptBytes, err := encrypt(key, []byte("pass1234"))
		check(err)
		MegaByte := &User{Email: "game1@cp-soft.com", Pass: b64.StdEncoding.EncodeToString(encryptBytes)}

		encryptBytes, err = encrypt(key, []byte("pass1235"))
		check(err)
		Stuffn := &User{Email: "game2@cp-soft.com", Pass: b64.StdEncoding.EncodeToString(encryptBytes)}

		encryptBytes, err = encrypt(key, []byte("pass1236"))
		check(err)
		Rooboy := &User{Email: "game3@cp-soft.com", Pass: b64.StdEncoding.EncodeToString(encryptBytes)}

		encryptBytes, err = encrypt(key, []byte("pass1237"))
		check(err)
		Shattered := &User{Email: "game4@cp-soft.com", Pass: b64.StdEncoding.EncodeToString(encryptBytes)}

		users["MegaByte"] = *MegaByte
		users["Stuffn"] = *Stuffn
		users["Rooboy"] = *Rooboy
		users["Shattered"] = *Shattered

		Sites["go2"] = users

		jsonString, _ := json.MarshalIndent(Sites, "", "  ")
		log.Println(string(jsonString))
		writeDataFile(jsonString)
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "clipboard"
	app.Usage = "Store all you passwords with me"

	//app.Flags = []cli.Flag{
	//	cli.StringFlag{
	//		Name:  "list, l",
	//		Value: "list",
	//		Usage: "List all sites and users",
	//	},
	//}

	app.Action = func(c *cli.Context) {
		action := "help"
		if len(c.Args()) > 0 {
			action = c.Args()[0]
		}
		switch action {
		case "add":
			add(c)
		case "get":
			get(c)
		case "list":
			list(c)
		case "help":
			log.Println("usage: clipboard")
			break
		default:
			log.Println("Unrecognized action - try `clipboard -h` for help.")
		}
	}
	app.Run(os.Args)
}

func add(c *cli.Context) {
	readDataFile()
	site := c.Args()[1]
	acct := c.Args()[2]
	email := c.Args()[3]
	pass := c.Args()[4]
	if Sites[site] == nil {
		newSite := make(map[string]User, 0)
		Sites[site] = newSite
	}

	encryptBytes, err := encrypt(key, []byte(pass))
	check(err)
	newUser := User{Email: email, Pass: string(b64.StdEncoding.EncodeToString(encryptBytes))}

	Sites[site][acct] = newUser
	jsonString, _ := json.MarshalIndent(Sites, "", "  ")
	log.Println(string(jsonString))
	writeDataFile(jsonString)
}

func get(c *cli.Context) {
	readDataFile()
	site := c.Args()[1]
	acct := c.Args()[2]
	data, _ := b64.StdEncoding.DecodeString(Sites[site][acct].Pass)
	pass, err := decrypt(key, []byte(data))
	check(err)
	err = clipboard.WriteAll(string(pass))
	check(err)
	fmt.Printf("Password for %s/%s copied to clipboad\n", site, acct)
}

func list(c *cli.Context) {
	readDataFile()
	for site, data := range Sites {
		fmt.Println("Site:", site)
		for user, info := range data {
			data, _ := b64.StdEncoding.DecodeString(info.Pass)
			pass, err := decrypt(key, []byte(data))
			check(err)
			fmt.Printf("\t%s: Email - %s, Password - %s\n", user, info.Email, pass)
		}
	}
}

func readDataFile() {
	dat, err := ioutil.ReadFile("data/clipboard.json")
	check(err)
	err = json.Unmarshal(dat, &Sites)
	check(err)
}

func writeDataFile(jsonString []byte) {
	err := ioutil.WriteFile("data/clipboard.json", jsonString, 0644)
	check(err)
}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
