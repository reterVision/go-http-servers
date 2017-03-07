package gin

import (
	"net/http"
	"testing"
)

func TestServe(t *testing.T) {
	addr := "127.0.0.1:8888"
	go Serve(addr)

	resp, err := http.Get("http://" + addr + "/hello-world")
	if err != nil {
		t.Errorf("Failed to get response: %s", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code != 200")
		return
	}
}
