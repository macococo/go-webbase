package net

import (
	"github.com/macococo/go-webbase/io"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, response interface{}) {
	WriteJsonBytesResponse(w, io.ToJsonBytes(response))
}

func WriteJsonBytesResponse(w http.ResponseWriter, response []byte) {
	w.Header()["Content-Type"] = []string{"application/json;charset=UTF-8"}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
