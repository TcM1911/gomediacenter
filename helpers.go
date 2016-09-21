package gomediacenter

import (
	"log"
	"net/http"
)

// CloseRespBody closes the response body and logs the error if it fails.
func CloseRespBody(r *http.Response) {
	if err := r.Body.Close(); err != nil {
		log.Printf("Error when closing the response body: %s", err.Error())
	}
}

// CloseReqBody closes the request body and logs the error if it fails.
func CloseReqBody(r *http.Request) {
	if err := r.Body.Close(); err != nil {
		log.Printf("Error when closing the request body: %s", err.Error())
	}
}
