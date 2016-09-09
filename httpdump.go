package main

import (
	"github.com/fatih/color"
	"fmt"
	"net/http"
	"net/http/httputil"
	"log"
	"time"
)

func dumperHandler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	t := time.Now()
	d := color.New(color.FgGreen, color.Bold)
	d.Printf("------------%s------------\n", t)
	fmt.Printf("%s\n\n", dump)
}

func main() {
	s := &http.Server{
		Addr:           ":8083",
		Handler:        http.HandlerFunc(dumperHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
