package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const (
	AIsInteger   = 1
	BIsNecessary = 2
)

func main() {
	http.HandleFunc("/test", testHandler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	resp := map[string]string{
		"error_code":    "0",
		"error_message": "success",
		"reference":     "111",
	}
	defer func() {
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatal("Error happened in JSON marshal ", err)
		}
		w.Write(jsonResp)
	}()

	a := req.URL.Query().Get("a")
	if a != "" {
		_, err := strconv.Atoi(a)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp["error_code"] = "1"
			resp["error_message"] = "parameter a should be an integer"
			return
		}
	}
	b := req.URL.Query().Get("b")
	if b == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp["error_code"] = "2"
		resp["error_message"] = "parameter b is necessary"
		// fmt.Fprintf(w, "bad request\n")
		return
	}
}
