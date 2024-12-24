package patterns

import "fmt"

type Handler interface {
	execute()
	setNext(Handler)
}

type Handler1 struct {
	next Handler
}

func (h *Handler1) execute() {
	fmt.Println("Executing on 1st handler")
	if h.next != nil {
		h.next.execute()
	}
}

func (h *Handler1) setNext(next Handler) {
	h.next = next
}

type Handler2 struct {
	next Handler
}

func (h *Handler2) execute() {
	fmt.Println("Executing on 2nd handler")
	if h.next != nil {
		h.next.execute()
	}
}

func (h *Handler2) setNext(next Handler) {
	h.next = next
}

func handler() {
	h1 := &Handler1{}
	h2 := &Handler2{}
	h1.setNext(h2)
	h1.execute()
}
