package common

import (
	"os"
	"os/signal"
	"syscall"
)

var signalsToCapture = []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}

// interruptHandler wraps systems signal receiving channel and provides a better interface for user
type interruptHandler struct {
	interruptChan chan os.Signal
}

// Wait blocks and waits for predefined system signal
func (ih *interruptHandler) Wait() {
	<-ih.interruptChan
	signal.Stop(ih.interruptChan)
}

// NewInterruptHandler captures system signals
func NewInterruptHandler() InterruptHandler {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, signalsToCapture...)

	return &interruptHandler{
		interruptChan: interruptChan,
	}
}

// SendSignalToCurrentProcess sends a signal to current process
func SendSignalToCurrentProcess(signal os.Signal) error {
	pid := os.Getpid()
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	process.Signal(signal)
	return nil
}
