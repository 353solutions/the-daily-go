package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/att/todo"
)

func rootDir() (string, error) {
	root := os.Getenv("TODO_ROOT")
	if len(root) > 0 {
		return root, nil
	}

	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return path.Join(user.HomeDir, ".todo"), nil
}

func main() {

}
