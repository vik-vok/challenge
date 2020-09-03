package main

import (
	"context"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"

	"log"
	"os"
)

func main() {
	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", p.ChallengeCreate); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}

	// Use PORT environment variable, or default to 8080.
	port := "8081"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
