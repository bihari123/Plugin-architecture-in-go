//go:build mage
// +build mage

package main

import (
	"errors"
	"fmt"
	"os/exec"
)

func BuildPlugin() error {
	cmd := exec.Command(
		"go",
		"build",
		"-buildmode=plugin",
		"-o",
		"plugin/myplugin.so",
		"plugin/myplugin.go",
	)

	if err := cmd.Run(); err != nil {
		// return errors.new(fmt.Printf("Error while building the plugin %v", err))
		return errors.New(fmt.Sprintf("Error while building the plugin %v", err.Error()))
	}
	return nil
}

func Run() error {
	cmd := exec.Command("go", "run", "cmd/main.go")

	if err := cmd.Run(); err != nil {
		// return errors.new(fmt.Printf("Error while building the plugin %v", err))
		return errors.New(fmt.Sprintf("Error while running the main program %v", err.Error()))
	}
	return nil
}
