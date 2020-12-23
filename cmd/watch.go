// Package cmd contains all the commands logic
// Copyright © 2020 Gld3n gld3ndev30@hotmail.com
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
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		w, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatalln(err)
		}
		defer w.Close()

		done := make(chan bool)

		if err = w.Add(args[0]); err != nil {
			log.Fatalln(err)
		}

		go watch(w)

		<-done
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ended running.")
	},
}

// watch will look for changes in the specified file.
func watch(w *fsnotify.Watcher) {
	fmt.Println("Started watching for changes.\n")
	for {
		select {
		case ev := <-w.Events:
			fmt.Printf("- Modified file: %s.\n", ev.Name)
			if ev.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("Starting the build process...\n")

			}

		case err := <-w.Errors:
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
