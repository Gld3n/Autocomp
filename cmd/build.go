// Copyright © 2020 Gld3m gld3ndev30@hotmail.com
package cmd

import (
	"fmt"
	fs "github.com/fsnotify/fsnotify"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Start the auto-compiling process",
	Run: func(cmd *cobra.Command, filename []string) {
		fmt.Println("build initialized...")

		//file := fmt.Sprint(filename[0:1])

		c := exec.Command("go", "run", "watcher.go")
		/*if err := c.Run(); err != nil {
			log.Fatal(err)
		}*/
		stdoutStderr, err := c.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", stdoutStderr)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func watcher() {
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

