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
		// get the first argument
		f := filename[0]

		fmt.Printf("Started watching at %s\n", f)
		// Establish a watcher for the file we want to observe for changes
		watch(f)
	},
}

// all the logic behind changing file monitoring
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
					//return
				}
				if event.Op&fs.Write == fs.Write {
					func() {
						fmt.Println("build initialized...")

						// specify the file we want to look at and the command
						cmd := exec.Command("go", "build", filename)

						// start compiling the file calling 'go build' and look for errors
						stdoutStderr, err := cmd.CombinedOutput()
						if err != nil {
							log.Fatal(err)
						}
						fmt.Printf("%s\n", stdoutStderr)

						fmt.Printf("File %s built succesfully\n", filename)

					}()
				}
			// return the function if not ok or log the error whether there's any
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// start watching the file specified
	err = watcher.Add(filename)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

