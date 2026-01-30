package handlers

import (
	"io"
	"log"
	"net/http"
)

type handlerObject struct {
	l *log.Logger
}

func NewHandler(l *log.Logger) *handlerObject {
	return &handlerObject{l}
}

func (h *handlerObject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("request made")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "bad stuff happened", http.StatusBadRequest)
	}
	log.Printf("Data is: %s\n", data)

	//---- write back to the user ----

	_, err = rw.Write([]byte("this is sent back to the client\n"))
	if err != nil {
		return
	}
}
