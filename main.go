/*
Command format: watcher [callback].
*/
package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var args = strings.Split(os.Args[1], " ")

// Listen and serve the file system file change events, and execute the callback command through child process when the file change events occur.
func main() {
	var watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()
	watcher.Add("./")
	var process *os.Process
	var path, _ = exec.LookPath(args[0])
	for event := range watcher.Events {
		if event.Op == fsnotify.Chmod {
			continue
		}
		go func() {
			if process != nil {
				process.Kill()
			}
			process, _ = os.StartProcess(path, args, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
			process.Wait()
		}()
	}
}
