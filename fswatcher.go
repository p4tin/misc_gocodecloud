package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func ExampleNewWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("Created file:", event, event.Name)
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event, event.Name)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Println("renamed file:", event, event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Add("./yaml")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {

	ExampleNewWatcher()
}
