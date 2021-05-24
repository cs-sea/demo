package main

import (
	"context"
	"demo/server/discover"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	handle := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}

	srv := &http.Server{
		Addr: "localhost:2223",
	}

	http.HandleFunc("/hello", handle)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatalln("shut down err", err)
		}
	}()

	dis := discover.NewConsulService()
	fmt.Println(dis.Register("hello", "localhost", 2223))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

}
