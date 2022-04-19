package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("call api ============")
		body := cekResi()
		w.Write(body)
	})

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)

	done := make(chan bool, 1)
	defer close(done)

	go func(done chan bool) {
		err := http.ListenAndServe(address, nil)
		if err != nil {
			fmt.Println(err.Error())
		}
		done <- true
	}(done)

	quit := make(chan os.Signal, 1)
	defer close(quit)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func(done chan bool) {
		fmt.Println("terima signal", <-quit)
		done <- true
	}(done)

	<-done
	fmt.Println("close mock-web")
}
