package main

import (
	"github.com/gin-gonic/gin"
	"ratings/routes"
	"ratings/tracer"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"log"
	"context"
)

func main() {

	tp, err := tracer.InitTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()


	router := gin.New()
	router.Use(otelgin.Middleware(tracer.ServiceName))
	routes.SetupRoute(router)
	router.Run(":8080")
}
