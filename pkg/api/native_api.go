package api

import (
	"fmt"
	"net/http"
)

type NativeServer struct {
	Port string
}

func (server *NativeServer) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(`{"message": "Welcome to VIZIX"}`))
}

func (server *NativeServer) Start()  {
	address := fmt.Sprintf(":%s", server.Port)
	http.Handle("/", server)
	http.ListenAndServe(address, nil)
}