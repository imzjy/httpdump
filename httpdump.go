package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"net/http"
	"net/http/httputil"
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

	fmt.Fprintf(w, "ok")
}

func main() {
	port := flag.Int("p", 8083, "listening port")
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)
	fmt.Println(addr)
	s := &http.Server{
		Addr:           addr,
		Handler:        http.HandlerFunc(dumperHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("server listening on port of %d\n", *port)
	log.Fatal(s.ListenAndServe())
}
