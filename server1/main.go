package main

import (
	"io"
	"log"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {
	getS2Command := "GetServer2"
	hystrix.ConfigureCommand(getS2Command, hystrix.CommandConfig{
		Timeout:                100,
		MaxConcurrentRequests:  5,
		RequestVolumeThreshold: 1,
		SleepWindow:            100,
		ErrorPercentThreshold:  1,
	})

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		body := "Hello, server1"
		err := hystrix.Do(getS2Command, func() error {
			if _, err := http.Get("http://localhost:3001"); err != nil {
				return err
			}

			body += " & server2"

			return nil
		}, nil)
		if err != nil {
			log.Println("cb err:", err)
		}

		io.WriteString(w, body+"!\n")
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
