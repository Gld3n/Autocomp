// Package cmd contains all the commands logic
// Copyright © 2020 Gld3n gld3ndev30@hotmail.com
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"time"

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

		fmt.Println("Started watching for changes.")

		// Add the directory of the given file instead of the file directly.
		for _, arg := range args {
			if err = w.Add(filepath.Dir(arg)); err != nil {
				log.Fatalln(err)
			}
		}

		// A new ticker to batch the incoming events.
		tick := time.Tick(300 * time.Millisecond)
		evs := make([]fsnotify.Event, 1)
		for {
			select {
			case ev, ok := <-w.Events:
				if !ok {
					return
				}

				if !inArgs(args, ev.Name) {
					continue
				}

				fmt.Printf("\n- Modified file: %s (%s).\n", ev.Name, ev.Op)

				// Add current event to array for batching.
				evs = append(evs, ev)
			case <-tick:
				// Checks on set interval if there are events.
				if len(evs) == 0 {
					continue
				}

				// Display messages for each event in batch.
				for _, event := range evs {
					if event.Op == fsnotify.Write {
						fmt.Printf("\nFile write detected: %v\n", event.Name)
					}
					if event.Op == fsnotify.Remove {
						fmt.Printf("\nFile delete detected: %v\n", event.Name)
					}
					if event.Op == fsnotify.Rename {
						fmt.Printf("\nFile rename detected: %v\n", event.Name)
					}

					// initialize the build process.
					err := build(args[0])
					if err != nil {
						log.Fatalln(err)
						return
					}
				}

				// Empty the batch array.
\				evs = make([]fsnotify.Event, 0)
			// Checksfor any error while the watching process.
			case err, ok := <-w.Errors:
				if !ok {
					return
				}

				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	},
}

// build is going to receive the file being watched as argument and compile it
// everytime an event is fired.
func build(file string) return error {
	cmd := exec.Command("go", "build", file)
	// Start running the command.
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	return nil
}

// Checks if the detected file is the one given as argument.
func inArgs(args []string, evName string) bool {
	for _, arg := range args {
		if evName == arg {
			return true
		}
	}

	// No matches as default
	return false
}
