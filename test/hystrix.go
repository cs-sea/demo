package main

import (
	"fmt"
	"html"
	"net"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {
	hystrix.ConfigureCommand("/test", hystrix.CommandConfig{
		Timeout:               1000,
		MaxConcurrentRequests: 1,
		ErrorPercentThreshold: 25,
	})

	hystrix.Go("/test", func() error {
		fmt.Println(1)
		return nil

	}, func(err error) error {
		fmt.Println("------------------")
		fmt.Println(err)
		return err
	})

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe("localhost:13333", nil)

}
