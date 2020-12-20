// Package cmd contains all the commands logic
// Copyright © 2020 Gld3m gld3ndev30@hotmail.com
package cmd

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command.
var watchCmd = &cobra.Command{
	Use:     "watch [--run] <file>",
	Aliases: []string{"w"},
	Short:   "Start the auto-compiling process",
	Example: "watch --run main.go",
	Run: func(cmd *cobra.Command, args []string) {
		w, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer w.Close()

		done := make(chan bool)
		go watch(w)

		err = w.Add("./main.go")
		if err != nil {
			log.Fatal(err)
		}
		<-done
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ended running.")
	},
}

// watch will look for changes in the specified file.
func watch(w *fsnotify.Watcher) {
	fmt.Println("Started watching for file changes.")
	for {
		select {
		case ev := <-w.Events:
			fmt.Printf("Modified file: %s.\n", ev.Name)
			if ev.Op&fsnotify.Write == fsnotify.Write {
				log.Println("Change detected.")
				fmt.Println("Starting the build process...")
			}

		case err := <-w.Errors:
			log.Fatal(err)
		}
	}
}
