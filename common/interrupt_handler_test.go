package common

import (
	"os"
	"testing"
)

func TestInterruptHandler(t *testing.T) {
	t.Parallel()

	for _, signal := range signalsToCapture {

		interruptHandler := NewInterruptHandler()

		go func(signal os.Signal) {
			err := SendSignalToCurrentProcess(signal)
			if err != nil {
				t.Errorf("Failed to Getpid: %s", err)
			}
		}(signal)

		interruptHandler.Wait()
	}
}
