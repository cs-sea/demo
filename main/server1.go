package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
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

	//e.Use(apmechov4.Middleware())
	go func() {
		e.Start("localhost:3031")
	}()
	e.GET("/", func(ctx echo.Context) error {

		fmt.Printf("%+v\n", ctx.Request())
		carrier := opentracing.HTTPHeadersCarrier(ctx.Request().Header)

		t := apmot.New()
		opentracing.SetGlobalTracer(t)
		parentSpan, _ := t.Extract(opentracing.HTTPHeaders, carrier)

		childSpan := t.StartSpan("ddd", opentracing.ChildOf(parentSpan))
		childSpan.SetTag("parent2", "parent2")
		childSpan.LogKV("event", "soft error",
			"type", "cache timeout",
			"waited.millis", 1500)
		defer childSpan.Finish()

		c := ctx.Request().Context()
		c = opentracing.ContextWithSpan(c, childSpan)

		c1, c := opentracing.StartSpanFromContext(c, "dd2")
		c1.LogKV("sss", "2222", "sss1", "5555")
		c1.LogKV(
			"event", "soft error",
			"type", "cache timeout",
			"waited.millis", 1500)
		c1.SetTag("tag1", 1011)
		time.Sleep(time.Second)
		c1.Finish()

		time.Sleep(time.Second)
		return ctx.JSON(http.StatusOK, 1)

	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	log.Println("Shutting down server...")
}
