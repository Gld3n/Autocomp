package cmd

import (
	"log"

	fs "github.com/fsnotify/fsnotify"
)

func Watcher() {
	watcher, err := fs.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fs.Write == fs.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("./build.go")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
