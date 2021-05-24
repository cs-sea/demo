package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
)

func main() {

	overseer.Run(overseer.Config{
		Program: prog,
		Address: ":3000",
		Fetcher: &fetcher.HTTP{
			URL:      "http://localhost:4000/binaries/myapp",
			Interval: 1 * time.Second,
		},
	})

	//// Start server
	//go func() {
	//	if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
	//		e.Logger.Fatal("shutting down the server")
	//	}
	//}()
	//
	//// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	//// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//if err := e.Shutdown(ctx); err != nil {
	//	e.Logger.Fatal(err)
	//}
}

//prog(state) runs in a child process
func prog(state overseer.State) {

	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "heh1e")
	})

	log.Printf("app (%s) listening...", state.ID)
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "app (%s) says hello\n", state.ID)
	}))
	http.Serve(state.Listener, e)
}
