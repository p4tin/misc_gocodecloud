package main

//export GODEBUG=cgocheck=0

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"io/ioutil"
	"os"

	b64 "encoding/base64"
	"github.com/atotto/clipboard"
)

var key []byte

type User struct {
	Email string
	Pass  string
}

var Sites map[string]map[string]User

func main() {
	content, err := ioutil.ReadFile("./aes.key")
	check(err)
	key, err = b64.StdEncoding.DecodeString(string(content))
	check(err)

	readDataFile()

	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Password Clipboard")
	window.Connect("destroy", gtk.MainQuit)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)

	table := gtk.NewTable(4, 10, false)

	x := uint(0)
	for site, data := range Sites {
		for user, info := range data {
			table.Attach(gtk.NewLabel(fmt.Sprintf("%s", site)), 0, 1, x, x+1, gtk.FILL, gtk.FILL, 4, 10)
			table.Attach(gtk.NewLabel(fmt.Sprintf("%s", user)), 1, 2, x, x+1, gtk.FILL, gtk.FILL, 4, 10)

			emailBtn := gtk.NewButtonWithLabel(fmt.Sprintf("%s", info.Email))
			emailBtn.Connect("clicked", func() {
				fmt.Println(emailBtn.GetLabel())
				err = clipboard.WriteAll(emailBtn.GetLabel())
				check(err)
			})
			table.Attach(emailBtn, 2, 3, x, x+1, gtk.FILL, gtk.FILL, 4, 10)
			data, _ := b64.StdEncoding.DecodeString(info.Pass)
			pass, err := decrypt(key, []byte(data))
			check(err)

			passBtn := gtk.NewButtonWithLabel(fmt.Sprintf("%s", pass))
			passBtn.Connect("clicked", func() {
				fmt.Println(passBtn.GetLabel())
				err = clipboard.WriteAll(passBtn.GetLabel())
				check(err)
			})
			table.Attach(passBtn, 3, 4, x, x+1, gtk.FILL, gtk.FILL, 4, 10)
			x++
		}
	}
	swin.AddWithViewPort(table)

	window.Add(swin)
	window.SetDefaultSize(450, 220)
	window.ShowAll()

	gtk.Main()
}

func readDataFile() {
	dat, err := ioutil.ReadFile("data/clipboard.json")
	check(err)
	err = json.Unmarshal(dat, &Sites)
	check(err)
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
	data, err := b64.StdEncoding.DecodeString(string(text))
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
