package main

import (
	playtone "cek-resi/play-tone"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {
	sigs := make(chan os.Signal, 1)
	defer close(sigs)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)
	defer close(done)

	go func(c chan bool) {
		for {
			checkResi(c)
			time.Sleep(10 * time.Minute)
		}
	}(done)

	go func() {
		fmt.Println("terima signal ", <-sigs)
		done <- true
	}()

	<-done
}

func checkResi(done chan bool) {
	count := 0

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("localhost"),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.OnHTML(`#resibelumlogin`, func(h *colly.HTMLElement) {
		count++
	})
	// Start scraping on http://coursera.com/browse
	c.Visit("http://localhost:9000/")
	fmt.Println("count =", count)

	if count > 6 {
		playtone.Play()
		done <- true
	}
}
