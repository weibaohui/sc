package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/weibaohui/sc/cmd"
)

func init() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		// Run Cleanup
		os.Exit(1)
	}()
}
func main() {

	cmd.Execute()
}
