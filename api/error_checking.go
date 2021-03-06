package main

import (
	"encoding/json"
	"net/http"
)

func databaseErrorRequest(w http.ResponseWriter, err error) bool {
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return true
	}
	return false
}

func decoderError(w http.ResponseWriter, err error) bool {
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return true
	}
	return false
}

func checkEmpty(w http.ResponseWriter, length int) bool {
	if length == 0 {
		w.WriteHeader(204)
		return true
	}
	return false
}

func forbiddenAuth(w http.ResponseWriter) {
	w.WriteHeader(403)
	json.NewEncoder(w).Encode("What are you trying to accomplish?")
}

// func catchErrors(fn MyFancyFunc) http.HandlerFunc {
// 	return funhttp.HandlerFunc(writer http.HandlerFunc, request *http.Request) {
// 		if err := fn(writer, request); err != nil {
// 			// do something with the error, eg encode it or log it or whatever
// 		}
// 	}
// }

// type (
// 	HttpFunc    func(http.HandlerFunc, *http.Request)
// 	MyFancyFunc func(http.HandlerFunc, *http.Request) error
// )
