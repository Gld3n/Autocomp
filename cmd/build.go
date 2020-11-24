// Copyright © 2020 Gld3m gld3ndev30@hotmail.com
package cmd

import (
	"fmt"
	fs "github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Start the auto-compiling process",
	Run: func(cmd *cobra.Command, filename []string) {
		f := filename[0]

		fmt.Printf("Started watching at %s\n", f)
		// Establish a watcher for the file we want to observe for changes
		watch(f)
	},
}

//
func watch(filename string) {
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
				if event.Op&fs.Write == fs.Write {
					func() {
						fmt.Println("build initialized...")

						cmd := exec.Command("go", "build", filename)

						stdoutStderr, err := cmd.CombinedOutput()
						if err != nil {
							log.Fatal(err)
						}
						fmt.Printf("%s\n", stdoutStderr)

						fmt.Printf("File %s built succesfully\n", filename)

					}()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(filename)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

