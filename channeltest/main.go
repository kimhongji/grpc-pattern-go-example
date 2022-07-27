package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var  sigChan chan os.Signal

func signalHandler() {
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
}
func main() {
	signalHandler()
	<-sigChan
	fmt.Println("sig")
	os.Exit(0)

}
