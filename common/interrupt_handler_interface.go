package common

// InterruptHandler defines methods for handling interrupt
type InterruptHandler interface {
	// Wait blocks and waits for predefined system signal
	Wait()
}
