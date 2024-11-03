package main

import (
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

func main() {
	var w, _ = fsnotify.NewWatcher()
	defer w.Close()
	w.Add("./")
	for e := range w.Events {
		if e.Op == fsnotify.Chmod {
			continue
		}
		var cmd = exec.Command("sh", "-c", os.Args[1])
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
