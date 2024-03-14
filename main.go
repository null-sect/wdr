package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func relativePath(base, target string) string {
	// get absolute path
	base, err := filepath.Abs(base)
	if err != nil {
		return target
	}
	target, err = filepath.Abs(target)
	if err != nil {
		return target
	}

	if base == "/" {
		return target
	}
	rel, err := filepath.Rel(base, target)
	if err != nil {
		return target
	}
	return rel
}

func flagParser(flags []string) ([]string, []string) {
	// only return start with "-"
	var flagsList []string
	var argsList []string
	for _, arg := range flags {
		if len(arg) > 0 && arg[0] == '-' {
			flagsList = append(flagsList, arg)
		} else {
			argsList = append(argsList, arg)
		}
	}

	//remove duplicated flags
	flagsMap := make(map[string]bool)
	var uniqueFlagsList []string
	for _, flag := range flagsList {
		if !flagsMap[flag] {
			flagsMap[flag] = true
			uniqueFlagsList = append(uniqueFlagsList, flag)
		}
	}
	flagsList = uniqueFlagsList

	return flagsList, argsList
}

func sliceChecker(key string, slice []string) bool {
	for _, s := range slice {
		if s == key {
			return true
		}
	}
	return false

}

func main() {
	root := "/"
	path := "."

	flags, args := flagParser(os.Args[1:])

	if sliceChecker("help", args) {
		fmt.Println("Usage: wdr [root] <path>")
		fmt.Println("  root: root of git repository")
		fmt.Println("  path: path to the file or directory")
		fmt.Println("  -h, --help: show this message")
		os.Exit(0)
	}

	if len(args) == 0 {
		fmt.Println(relativePath(root, path))
		return
	} else if len(args) == 1 {
		path = args[0]
	} else if len(args) == 2 {
		root = args[0]
		path = args[1]
	} else {
		fmt.Println("invalid arguments")
		fmt.Println("please check the usage")
		fmt.Println("  wdr help")
		os.Exit(1)
	}
	// nno flags
	if len(flags) == 0 {
		fmt.Println(relativePath(root, path))
		return
	}

	if sliceChecker("-h", flags) || sliceChecker("--help", flags) {
		fmt.Println("Usage: wdr [root] <path>")
		fmt.Println("  root: root of git repository")
		fmt.Println("  path: path to the file or directory")
		fmt.Println("  -h, --help: show this message")
		os.Exit(0)
	} else if sliceChecker("-v", flags) || sliceChecker("--version", flags) {
		fmt.Println("wdr 0.1.0")
		os.Exit(0)
	}

	if !((sliceChecker("-r", flags) || sliceChecker("--root", flags)) && (sliceChecker("-p", flags) || sliceChecker("--path", flags))) {
		fmt.Println("invalid flags")
		fmt.Println("please check the usage")
		fmt.Println("  wdr help")
		os.Exit(1)
	} else if sliceChecker("-r", flags) || sliceChecker("--root", flags) {
		for i, flag := range flags {
			if flag == "-r" || flag == "--root" {
				root = os.Args[i]
				break
			}
		}
	} else if sliceChecker("-p", flags) || sliceChecker("--path", flags) {
		for i, flag := range flags {
			if flag == "-p" || flag == "--path" {
				path = os.Args[i]
				break
			}
		}
	} else {
		fmt.Println("invalid flags")
		fmt.Println("please check the usage")
		fmt.Println("  wdr help")
		os.Exit(1)
	}

	fmt.Println(relativePath(path, root))
}
