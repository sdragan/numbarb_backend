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

	if !ok || len(numbers[0]) < 1 {
		return
	}

	number, err := strconv.Atoi(numbers[0])
	if err == nil {
		fmt.Fprintf(w, converters.ConvertIntegerToWords(int(number)))
	}
}
