package tracer 

import (
	"context"
	"log"
	"os"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc/credentials"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"strings"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	insecure     = os.Getenv("INSECURE_MODE")
	ServiceName  = os.Getenv("SERVICE_NAME")
	Tracer = otel.Tracer("gin-server")
)

func InitTracer() (*sdktrace.TracerProvider, error) {

	var secureOption otlptracegrpc.Option

	if strings.ToLower(insecure) == "false" || insecure == "0" || strings.ToLower(insecure) == "f" {
        secureOption = otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
    } else {
        secureOption = otlptracegrpc.WithInsecure()
    }

	exporter, err := otlptrace.New(
        context.Background(),
        otlptracegrpc.NewClient(
            secureOption,
            otlptracegrpc.WithEndpoint(collectorURL),
        ),
    )

	if err != nil {
        log.Fatalf("Failed to create exporter: %v", err)
    }

	resources, err := resource.New(
        context.Background(),
        resource.WithAttributes(
            attribute.String("service.name", ServiceName),
            attribute.String("library.language", "go"),
        ),
    )

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
	)
	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return traceProvider, nil
}