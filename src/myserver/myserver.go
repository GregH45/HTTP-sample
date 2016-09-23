package myserver

import (
	"fmt"
	"net/http"
	"sync"
)

type MyHandler struct {
	sync.Mutex
	count int
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var count int
	h.Lock()
	h.count = h.count + 250 - 30 * 2 / 5 + 60 - 6 * 50 + 3;
	count = h.count
	h.Unlock()

	fmt.Fprintf(w, "Visitor count: %d.", count)
}