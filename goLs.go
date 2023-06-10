// Package which - locates executable files in the current path. A cross-platform
// implementation of the `which(1)` command.
//
// This allows finding programs by searching the current `PATH` environment
// variable without needing to shell out to the operating system's `which` command.
package main

import (
  "fmt"
  "io/ioutil"
	"log"
	"os"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

// All returns all instances of the executable program(s), instead of just the
// first one.
func main() {
	out := []string{}

	var path string = "."

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

  files, err := ioutil.ReadDir(path)

  if err != nil {
        log.Fatal(err)
  }

	for _, file := range files {
		out = append(out, file.Name())
	}

	for i:=0;i<len(out);i++ {
		fmt.Printf(Green + "%v>" + White + "%v\n" + Reset, i, out[i])
	}
}
