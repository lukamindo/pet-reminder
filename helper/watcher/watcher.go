package watcher

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"gopkg.in/fsnotify.v1"
)

const defaultPeriod = time.Second * 2

// Exported variables
var (
	// WatchFilename is a filename to watch for, by default it is a program binary
	WatchFilename string
	// WatchPeriod is a period of time to check for changes, default is `1 * time.Second
	WatchPeriod = defaultPeriod
	// RestartFunc is a function called for restart,
	// default is `RestartByExec` you can use `SendSIGUSR2` or your custom function
	RestartFunc = RestartByExec
)

var (
	executableArgs []string
	executableEnvs []string
	executablePath string
	ticker         *time.Ticker
	startFileInfo  os.FileInfo
	listeners      []chan bool
)

func init() {
	listeners = make([]chan bool, 0)
	executableArgs = os.Args
	executableEnvs = os.Environ()
	executablePath, _ = filepath.Abs(os.Args[0])
	WatchFilename = executablePath
}

// StartWatcher starts timer
func StartWatcher() {
	ticker = time.NewTicker(WatchPeriod)
	go watcher()
	go watchFileChanges()
}

// GetNotifier returns a channel, it will recived message before restart
// channel is synchronous and must be readed to continue
func GetNotifier() (c chan bool) {
	c = make(chan bool)
	listeners = append(listeners, c)
	return c
}

func watcher() {
	for range ticker.C {
		if isChanged() {
			notify()
			RestartFunc()
		}
	}
}

func isChanged() bool {
	return isChangedByStat()
}

func isChangedByStat() bool {
	fileinfo, err := os.Stat(WatchFilename)
	if err == nil {
		// first update
		if startFileInfo == nil {
			startFileInfo = fileinfo
			return false
		}
		// non-first update
		if startFileInfo.ModTime() != fileinfo.ModTime() ||
			startFileInfo.Size() != fileinfo.Size() {
			return true
		}

		return false
	}

	log.Printf("cannot find %s: %s", WatchFilename, err)
	return false
}

func notify() {
	for _, c := range listeners {
		c <- true
	}
}

// RestartByExec calls `syscall.Exec()` to restart app
func RestartByExec() {
	binary, err := exec.LookPath(executablePath)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
	time.Sleep(1 * time.Second)
	execErr := syscall.Exec(binary, executableArgs, executableEnvs)
	if execErr != nil {
		log.Printf("error: %s %v", binary, execErr)
	}
}

// Watch file changes and restart server
func watchFileChanges() {
	ctx := context.Background()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Panic(context.Background(), err)
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
				if event.Op&fsnotify.Write == fsnotify.Write {
					if filepath.Ext(event.Name) == ".go" {
						rebuild()
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Panic(ctx, err)
			}
		}
	}()
	basePath, err := os.Getwd()
	if err != nil {
		log.Panic(context.Background(), err)
	}
	err = filepath.Walk(basePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(info.Name()) == ".go" {
				err = watcher.Add(filepath.Dir(path))
				if err != nil {
					log.Panic(ctx, err)
				}
			}
			return nil
		},
	)
	if err != nil {
		log.Panic(ctx, err)
	}
	<-done
}

func rebuild() {
	cmd := exec.Command("go", "build", "-o", "bin/pet-reminder")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("rebuild err: ", err)
	}
}
