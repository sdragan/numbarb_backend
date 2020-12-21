package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sdragan/numbarb_backend/converters"
)

func main() {
	http.HandleFunc("/convert", convert)
	http.ListenAndServe(":3000", nil)
}

func convert(w http.ResponseWriter, req *http.Request) {
	numbers, ok := req.URL.Query()["number"]
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if !ok || len(numbers[0]) < 1 {
		fmt.Fprintf(w, "bad request or missing param")
		return
	}

	number, err := strconv.Atoi(numbers[0])
	if err == nil {
		fmt.Fprintf(w, converters.ConvertIntegerToWords(int(number)))
	} else {
		fmt.Fprintf(w, "couldn't get number")
	}
}
