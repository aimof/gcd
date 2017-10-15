package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		return
	}

	pwdChild, err := filepath.Abs(args[1])
	if err == nil {
		info, err := os.Stat(pwdChild)
		if err == nil {
			if info.IsDir() {
				fmt.Print(pwdChild)
				return
			}
		}
	}

	targetDirs := filepath.SplitList(args[1])
	if len(targetDirs) == 0 {
		return
	}
	gcdpath := os.Getenv("GCDPATH")
	if gcdpath == "" {
		return
	}
	paths := strings.Split(gcdpath, ":")

	for _, p := range paths {
		if p == "" {
			continue
		}
		absPath, err := filepath.Abs(p)
		if err != nil {
			continue
		}
		if targetDirs[0] == filepath.Base(absPath) {
			fmt.Print(absPath)
			return
		}

		absTarget := filepath.Join(absPath, args[1])
		info, err := os.Stat(absTarget)
		if err != nil {
			continue
		}
		if info.IsDir() {
			fmt.Print(absTarget)
			return
		}
	}
	fmt.Print(args[1])
	return
}
