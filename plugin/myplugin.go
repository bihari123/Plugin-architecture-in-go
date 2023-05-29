package main

import "fmt"

type MyPlugin struct{}

func (p MyPlugin) Run() {
	fmt.Println("Running Plugin...")
}

// exported

var NewPlugIn = MyPlugin{}
