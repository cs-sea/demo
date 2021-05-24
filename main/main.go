package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmot"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	fmt.Println(os.Getenv("ELASTIC_APM_SERVER_URL"))
	fmt.Println(os.Getenv("ELASTIC_APM_LOG_LEVEL"))

	e := echo.New()
	go func() {
		e.Start("localhost:3030")
	}()
	e.GET("/", func(ctx echo.Context) error {

		t1, _ := apm.NewTracer("i", "v1")
		t := apmot.New(apmot.WithTracer(t1))
		parnet, _ := opentracing.StartSpanFromContextWithTracer(ctx.Request().Context(), t, "test")
		defer parnet.Finish()

		request, _ := http.NewRequest(http.MethodGet, "http://localhost:3031", nil)

		_ = t.Inject(parnet.Context(), opentracing.HTTPHeaders, request.Header)

		time.Sleep(time.Second)
		_, _ = http.DefaultClient.Do(request)
		return ctx.JSON(http.StatusOK, 1)

	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	log.Println("Shutting down server...")
}
