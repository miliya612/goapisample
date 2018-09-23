package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Responder interface {
	Write(w http.ResponseWriter)
	Status() int
}

type Response struct {
	status int
	header http.Header
	body []byte
}

func (r *Response) Write(w http.ResponseWriter) {
	header := w.Header()
	for k, v := range r.header {
		header[k] = v
	}
	w.WriteHeader(r.status)
	w.Write(r.body)
}

func (r Response) Status() int {
	return r.status
}

func (r Response) Header(k, v string) *Response {
	r.header.Set(k, v)
	return &r
}

func Empty(status int) *Response {
	return Respond(status, nil)
}

func Json(status int, body interface{}) *Response {
	return Respond(status, body).Header("Content-Type", "application/json; charset=UTF-8")
}

func Created(status int, body interface{}, location string) *Response {
	return Json(status, body).Header("Location", location)
}

func Error(status int, message string, err error) *Response {
	log.Printf("[%d]\t%s:\t%s", status, message, err)
	switch status / 100 {
	case 4, 5:
		return Respond(status, message).Header("Content-Type", "application/json; charset=UTF-8")
	default:
		panic("status code is not 4xx or 5xx")
	}
}

func Respond(status int, body interface{}) *Response {
	var b []byte
	var err error
	switch t := body.(type) {
	case string:
		b = []byte(t)
	default:
		if b, err = json.Marshal(body); err != nil {
			return Error(http.StatusInternalServerError, "failed marshalling json", err)
		}
	}

	return &Response{
		status: status,
		body: b,
		header: make(http.Header),
	}
}