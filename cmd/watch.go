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

		for _, arg := range args {
			if err = w.Add(filepath.Dir(arg)); err != nil {
				log.Fatalln(err)
			}
		}

		ticker := time.NewTicker(300 * time.Millisecond)
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
			case <-ticker.C:
				// Checks on set interval if there are events.
				if len(evs) > 0 {
					// Display messages for each event in batch.
					for _, event := range evs {
						if event.Op == fsnotify.Write {
							fmt.Printf("\nFile write detected: %v (%v)\n", event.Name, event.Op)
						}
						if event.Op == fsnotify.Remove {
							fmt.Printf("\nFile delete detected: %v (%v)\n", event.Name, event.Op)
						}
						if event.Op == fsnotify.Rename {
							fmt.Printf("\nFile rename detected: %v (%v)\n", event.Name, event.Op)
						}

						build()
						// Empty the batch array.
						evs = make([]fsnotify.Event, 0)
					}
				} else {
					continue
				}

				evs = make([]fsnotify.Event, 0)
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

func build() {
	cmd := exec.Command("go", "build", "main.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

func inArgs(args []string, evName string) bool {
	for _, arg := range args {
		if evName == arg {
			return true
		}
	}

	return false
}
