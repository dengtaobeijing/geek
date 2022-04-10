package test

import (
	"encoding/json"
	"net/http"
)

type HelloDTO struct {
	Name     string
	UserName string
}
type UserHandler interface {
	SayHello(id string) *HelloDTO
}

func NewHelloHandler(user UserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		id := r.Form.Get("id")
		hello := user.SayHello(id)
		bytes, _ := json.Marshal(hello)
		_, _ = w.Write(bytes)
	}
}
