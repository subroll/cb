package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ticker := time.NewTicker(2000 * time.Millisecond)
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	go func() {
		for {
			select {
			case <-ticker.C:
				for i := 0; i < 200; i++ {
					go func() {
						resp, err := http.Get("http://localhost:3000")
						if err != nil {
							log.Println("resp err:", err)
							return
						}

						var body []byte
						if resp != nil {
							body, err = io.ReadAll(resp.Body)
							if err != nil {
								log.Println("read err:", err)
								return
							}
						}

						if body != nil {
							log.Println("body:", string(body))
						}
					}()
				}
			}
		}
	}()

	<-sigint
}
